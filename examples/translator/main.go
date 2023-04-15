package main

import (
	"context"
	"fmt"
	"github.com/pkg-id/goval"
	"github.com/pkg-id/goval/errtrans"
	"github.com/pkg-id/goval/govalregex"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func init() {
	bundle, err := errtrans.DefaultBundle()
	if err != nil {
		panic(err)
	}

	translator := errtrans.NewTranslator(errtrans.WithBundle(bundle))
	goval.SetErrorTranslator(translator)
}

func main() {
	usr := User{
		ID:    "",
		Name:  "b",
		Email: "bob-mail?",
	}

	ctx := context.Background()
	ctx = errtrans.ContextWithLanguage(ctx, "en")
	err := goval.Execute(ctx,
		goval.Named("id", usr.ID, goval.String().Required()),
		goval.Named("name", usr.Name, goval.String().Required().Min(3).Max(20)),
		goval.Named("email", usr.Email, goval.String().Required().Match(govalregex.AlphaNumeric)),
	)
	fmt.Println(err)
}
