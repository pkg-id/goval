package errtrans

import (
	"bytes"
	"context"
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/pkg-id/goval"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	ErrBundleIsNoSet                = goval.TextError("bundle is not set")
	ErrLanguageDictionaryIsNotFound = goval.TextError("language dictionary is not found")
	ErrLanguageKeyIsNotFound        = goval.TextError("language key is not found")
)

//go:embed locale
var localeFS embed.FS

type Bundle map[string]Dictionary
type Dictionary map[string]string

type contextType struct {
	name string
}

var languageContext = contextType{name: "language"}

func ContextWithLanguage(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, languageContext, lang)
}

func LanguageFromContext(ctx context.Context, fallback string) string {
	lang, ok := ctx.Value(languageContext).(string)
	if !ok {
		return fallback
	}
	return lang
}

var ruleCodeToTemplateKey = map[goval.RuleCoder]string{
	goval.StringRequired: "strings.required",
	goval.StringMin:      "strings.min",
	goval.StringMax:      "strings.max",
	goval.StringMatch:    "strings.match",
	goval.StringIn:       "strings.in",
	goval.StringInFold:   "strings.in_fold",
	goval.NumberRequired: "numbers.required",
	goval.NumberMin:      "numbers.min",
	goval.NumberMax:      "numbers.max",
	goval.SliceRequired:  "slices.required",
	goval.SliceMin:       "slices.min",
	goval.SliceMax:       "slices.max",
	goval.TimeRequired:   "times.required",
	goval.TimeMin:        "times.min",
	goval.TimeMax:        "times.max",
	goval.MapRequired:    "maps.required",
	goval.MapMin:         "maps.min",
	goval.MapMax:         "maps.max",
	goval.PtrRequired:    "pointers.required",
}

type Option func(t *Translator)

func WithBundle(bundle Bundle) Option {
	return func(t *Translator) { t.bundle = bundle }
}

func DefaultBundle() (Bundle, error) {
	bundle := make(Bundle)
	err := fs.WalkDir(localeFS, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// skip if entry is a directory
		if entry.IsDir() {
			return nil
		}

		// skip if the entry is not JSON file.
		ext := filepath.Ext(entry.Name())
		if ext != ".json" {
			return nil
		}

		file, err := localeFS.Open(path)
		if err != nil {
			return fmt.Errorf("open file. path=%s: %w", path, err)
		}

		dict := make(Dictionary)
		if err := json.NewDecoder(file).Decode(&dict); err != nil {
			return fmt.Errorf("decode json: %w", err)
		}

		if err := file.Close(); err != nil {
			return fmt.Errorf("closing the file: %w", err)
		}

		lang := strings.TrimSuffix(filepath.Base(entry.Name()), ext)
		bundle[lang] = dict
		return nil
	})

	return bundle, err
}

type Translator struct {
	tpl    *template.Template
	bundle Bundle
}

func NewTranslator(opts ...Option) goval.ErrorTranslator {
	t := Translator{
		tpl:    template.New("translator"),
		bundle: nil,
	}

	for _, opt := range opts {
		opt(&t)
	}

	return &t
}

func (t *Translator) translate(ctx context.Context, err *goval.RuleError, key string) error {
	if len(t.bundle) == 0 {
		return ErrBundleIsNoSet
	}

	dict, ok := t.bundle[LanguageFromContext(ctx, "en")]
	if !ok {
		return ErrLanguageDictionaryIsNotFound
	}

	textTemplate, ok := dict[key]
	if !ok {
		return ErrLanguageKeyIsNotFound
	}

	tpl, errParse := t.tpl.Parse(textTemplate)
	if errParse != nil {
		return goval.TextError(errParse.Error())
	}

	buff := bytes.Buffer{}
	if err := tpl.Execute(&buff, err); err != nil {
		return goval.TextError(err.Error())
	}

	return goval.TextError(buff.String())
}

func (t *Translator) Translate(ctx context.Context, err *goval.RuleError) error {
	key, ok := ruleCodeToTemplateKey[err.Code]
	if !ok {
		return goval.TextError(fmt.Sprintf("RuleError[code=%v] is not registered yet.", err.Code))
	}
	return t.translate(ctx, err, key)
}
