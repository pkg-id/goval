package govalregex

import (
	"regexp"
	"sync"
)

var caches = sync.Map{}

type expr string

func (e expr) String() string { return string(e) }
func (e expr) RegExp() *regexp.Regexp {
	cached, ok := caches.Load(e)
	if ok {
		return cached.(*regexp.Regexp)
	}

	re := regexp.MustCompile(string(e))
	caches.Store(e, re)
	return re
}

// These regex expression are copied from https://github.com/go-playground/validator/blob/master/regexes.go.
const (
	Alpha                 = expr("^[a-zA-Z]+$")
	AlphaNumeric          = expr("^[a-zA-Z0-9]+$")
	AlphaUnicode          = expr("^[\\p{L}]+$")
	AlphaUnicodeNumeric   = expr("^[\\p{L}\\p{N}]+$")
	Numeric               = expr("^[-+]?[0-9]+(?:\\.[0-9]+)?$")
	Number                = expr("^[0-9]+$")
	Hexadecimal           = expr("^(0[xX])?[0-9a-fA-F]+$")
	HexColor              = expr("^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$")
	RGB                   = expr("^rgb\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*\\)$")
	RGBA                  = expr("^rgba\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$")
	HSL                   = expr("^hsl\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*\\)$")
	HSLA                  = expr("^hsla\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$")
	Email                 = expr("^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
	E164                  = expr("^\\+[1-9]?[0-9]{7,14}$")
	Base64                = expr("^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$")
	Base64URL             = expr("^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$")
	Base64RawURL          = expr("^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2,4})$")
	ISBN10                = expr("^(?:[0-9]{9}X|[0-9]{10})$")
	ISBN13                = expr("^(?:(?:97(?:8|9))[0-9]{10})$")
	UUID3                 = expr("^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$")
	UUID4                 = expr("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")
	UUID5                 = expr("^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")
	UUID                  = expr("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	UUID3RFC4122          = expr("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")
	UUID4RFC4122          = expr("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
	UUID5RFC4122          = expr("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$")
	UUIDRFC4122           = expr("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")
	ULID                  = expr("^[A-HJKMNP-TV-Z0-9]{26}$")
	MD4                   = expr("^[0-9a-f]{32}$")
	MD5                   = expr("^[0-9a-f]{32}$")
	SHA256                = expr("^[0-9a-f]{64}$")
	SHA384                = expr("^[0-9a-f]{96}$")
	SHA512                = expr("^[0-9a-f]{128}$")
	RIPEMD128             = expr("^[0-9a-f]{32}$")
	RIPEMD160             = expr("^[0-9a-f]{40}$")
	Tiger128              = expr("^[0-9a-f]{32}$")
	Tiger160              = expr("^[0-9a-f]{40}$")
	Tiger192              = expr("^[0-9a-f]{48}$")
	ASCII                 = expr("^[\x00-\x7F]*$")
	PrintableASCII        = expr("^[\x20-\x7E]*$")
	Multibyte             = expr("[^\x00-\x7F]")
	DataURI               = expr(`^data:((?:\w+\/(?:([^;]|;[^;]).)+)?)`)
	Latitude              = expr("^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$")
	Longitude             = expr("^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$")
	SSN                   = expr(`^[0-9]{3}[ -]?(0[1-9]|[1-9][0-9])[ -]?([1-9][0-9]{3}|[0-9][1-9][0-9]{2}|[0-9]{2}[1-9][0-9]|[0-9]{3}[1-9])$`)
	HostnameRFC952        = expr(`^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`)                                                                   // https://tools.ietf.org/html/rfc952)
	HostnameRFC1123       = expr(`^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62}){1}(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?$`)                                 // accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123)
	FQDNRFC1123           = expr(`^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$`) // same as hostnameRFC1123 but must contain a non-numerical TLD (possibly ending with '.'
	BTCAddress            = expr(`^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$`)                                                                             // bitcoin address)
	BTCAddressUpperBech32 = expr(`^BC1[02-9AC-HJ-NP-Z]{7,76}$`)                                                                                   // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	BTCAddressLowerBech32 = expr(`^bc1[02-9ac-hj-np-z]{7,76}$`)                                                                                   // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	ETHAddress            = expr(`^0x[0-9a-fA-F]{40}$`)
	ETHAddressUpper       = expr(`^0x[0-9A-F]{40}$`)
	ETHAddressLower       = expr(`^0x[0-9a-f]{40}$`)
	URLEncoded            = expr(`^(?:[^%]|%[0-9A-Fa-f]{2})*$`)
	HTMLEncoded           = expr(`&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?`)
	HTML                  = expr(`<[/]?([a-zA-Z]+).*?>`)
	JWT                   = expr("^[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]*$")
	SplitParams           = expr(`'[^']*'|\S+`)
	Bic                   = expr(`^[A-Za-z]{6}[A-Za-z0-9]{2}([A-Za-z0-9]{3})?$`)
	Semver                = expr(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`) // numbered capture groups https://semver.org/
	DNSRFC1035Label       = expr("^[a-z]([-a-z0-9]*[a-z0-9]){0,62}$")
	CVE                   = expr(`^CVE-(1999|2\d{3})-(0[^0]\d{2}|0\d[^0]\d{1}|0\d{2}[^0]|[1-9]{1}\d{3,})$`) // CVE Format Id https://cve.mitre.org/cve/identifiers/syntaxchange.html
	Mongodb               = expr("^[a-f\\d]{24}$")
	Cron                  = expr(`(@(annually|yearly|monthly|weekly|daily|hourly|reboot))|(@every (\d+(ns|us|µs|ms|s|m|h))+)|((((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|\*) ?){5,7})`)
)
