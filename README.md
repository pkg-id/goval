# Goval
[![GoDoc](https://godoc.org/github.com/pkg-id/env?status.svg)](https://godoc.org/github.com/pkg-id/goval)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/pkg-id/goval/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/pkg-id/goval)](https://goreportcard.com/report/github.com/pkg-id/goval)
[![codecov](https://codecov.io/gh/pkg-id/goval/branch/develop/graph/badge.svg?token=FCC0ZNH1D7)](https://codecov.io/gh/pkg-id/goval)

**Goval** or **Go Validator** is a package for value validation in Go.

This Go packages aims to provide a user-friendly validation library that is easy to integrate, extend, and use. 
It utilizes function composition for building complex validation logic and avoids reflection for improved performance. 
It is designed to enhance the capabilities of Go functions and is safe for concurrent use.

## Features

* [Reusable Validation Rules](#reusable-validation-rules)
* [Extendable Validation Rules](#extendable-validation-rules)
* [Customizable Validation Rules](#customizable-validation-rules)
* [Composable Validation Rules](#composable-validation-rules)
* [No Reflection](#no-reflection)
* [Support for Customizable and Translatable Error Messages](#support-for-customizable-and-translatable-error-messages)

## Installation

```shell
go get github.com/pkg-id/goval
```

## Feature Details

### Reusable Validation Rules

This means that you can define the validation rules once and use them many times for different values with the same type. For example:

```go
validator := goval.String().Required().Min(2).Max(9)

ctx := context.Background()
fmt.Println(validator.Validate(ctx, ""))           // err: {"code":2000}
fmt.Println(validator.Validate(ctx, "h"))          // err: {"code":2001,"args":[2]}
fmt.Println(validator.Validate(ctx, "0123456789")) // err: {"code":2002,"args":[9]}
```

The `validator` function is used to validate strings with values `""`, `"h"`, and `"0123456789"`.

### Extendable Validation Rules

This means that if you already have common validation rules, you can create a new one from them without modifying the existing behavior. For example:

```go
validator := goval.String().Required().Min(2).Max(9)
extendedValidator := validator.Match(govalregex.AlphaNumeric)

ctx := context.Background()
fmt.Println(validator.Validate(ctx, "hello!"))          // err: <nil>
fmt.Println(extendedValidator.Validate(ctx, "hello!"))  // err: {"code":2003,"args":["^[a-zA-Z0-9]+$"]}
```

Both `validator` and `extendedValidator` validate the same input `"hello!"`.
The original `validator` (or the parent) will be valid, since it does not have rules for checking alphanumeric. But the `extendedValidator` is not valid.

### Customizable Validation Rules

This means that you can define your own validation rules and use them along with the predefined rules. For example, we can define a rule to check if a given string has a prefix that we want. First, let's create the validation rule as follows:

```go
func HasPrefix(prefix string) goval.StringValidator {
	return func(ctx context.Context, value string) error {
		if !strings.HasPrefix(value, prefix) {
			return goval.NewRuleError(ECHasPrefix, prefix) 
		}
		return nil
	}
}
```

The `HasPrefix` function will check if the input value is prefixed with the given `prefix`. If not, it will return an error. Every error in goval is expected to have an error code, which is useful for generating human-readable messages. To create a new error code, we can implement the `goval.RuleCoder` interface as shown below:

```go
type MyCustomErrorCode string

const (
	ECHasPrefix = MyCustomErrorCode("ec-has-prefix")
)

func (e MyCustomErrorCode) Equal(other goval.RuleCoder) bool {
	val, ok := other.(MyCustomErrorCode)
	return ok && e == val
}
```

Finally, we can chain our custom validation rule by using the With method as shown below:

```go
validator := goval.String().Required().Min(2).Max(9).With(HasPrefix(":"))

ctx := context.Background()
fmt.Println(validator.Validate(ctx, "abc")) // err: {"code":"ec-has-prefix","args":[":"]}
```

This will create a new validator that includes our custom rule, and will validate strings that meet all the defined criteria, including having the specified `prefix`.

### Composable Validation Rules

As we saw previously, we only used a single rules chain, which is boring! Most of the time, we deal with struct, map, or slice rather than a single value. This package is also aware of that. Let's take the following struct as an example:

```go
type SocialMedia struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type RegisterRequest struct {
	Name            string        `json:"name"`
	Phone           string        `json:"phone"`
	Age             uint          `json:"age"`
	Height          float64       `json:"height"`
	SocialMediaList []SocialMedia `json:"social_media_list"`
}
```

For example, we got data like the following:

```go
req := RegisterRequest{
    Name:   "",
    Phone:  "",
    Age:    16,
    Height: 172.5,
    SocialMediaList: []SocialMedia{
        {Name: "", Link: ""},
    },
}
```
We want to combine all fields into a group and expect a single error as the result if any rules are violated. To do that, we can use `goval.Execute` to group all validator rules, and use `goval.Bind` to bind the data to the validator.


```go
ctx := context.Background()
err := goval.Execute(ctx,
    goval.Bind[string](req.Name, goval.String().Required().Min(2).Max(20)),
    goval.Bind[string](req.Phone, goval.String().Required().Match(govalregex.E164)),
    goval.Bind[uint](req.Age, goval.Number[uint]().Required().Min(17)),
    goval.Bind[float64](req.Height, goval.Number[float64]().Required()),
    goval.Bind[[]SocialMedia](req.SocialMediaList, goval.Slice[SocialMedia]().Required().EachFunc(SocialMediaValidator)),
)
fmt.Println(err)
```

The `SocialMediaValidator` is a custom validation:

```go
func SocialMediaValidator(ctx context.Context, s SocialMedia) error {
	return goval.Execute(ctx,
		goval.Bind[string](s.Name, goval.String().Required()),
		goval.Bind[string](s.Link, goval.String().Required()),
	)
}
```

And the error result will look like the following:

```json
[
  {
    "code":2000
  },
  {
    "code":2000
  },
  {
    "code":3001,
    "args":[
      17
    ]
  },
  [
    [
      {
        "code":2000
      },
      {
        "code":2000
      }
    ]
  ]
]
```

The error structure is correct, but we need to add field names. To do that, we just need to change `goval.Bind` to` gobal.Named` as follows:


```go
func SocialMediaValidator(ctx context.Context, s SocialMedia) error {
	return goval.Execute(ctx,
		goval.Named[string]("name", s.Name, goval.String().Required()),
		goval.Named[string]("link", s.Link, goval.String().Required()),
	)
}
```

```go
ctx := context.Background()
err := goval.Execute(ctx,
    goval.Named[string]("name", req.Name, goval.String().Required().Min(2).Max(20)),
    goval.Named[string]("phone", req.Phone, goval.String().Required().Match(govalregex.E164)),
    goval.Named[uint]("age", req.Age, goval.Number[uint]().Required().Min(17)),
    goval.Named[float64]("height", req.Height, goval.Number[float64]().Required()),
    goval.Named[[]SocialMedia]("social_media_list", req.SocialMediaList, goval.Slice[SocialMedia]().Required().EachFunc(SocialMediaValidator)),
)
```

And the mew error will be look like this:

```json
[
  {
    "key":"name",
    "err":{
      "code":2000
    }
  },
  {
    "key":"phone",
    "err":{
      "code":2000
    }
  },
  {
    "key":"age",
    "err":{
      "code":3001,
      "args":[
        17
      ]
    }
  },
  {
    "key":"social_media_list",
    "err":[
      [
        {
          "key":"name",
          "err":{
            "code":2000
          }
        },
        {
          "key":"link",
          "err":{
            "code":2000
          }
        }
      ]
    ]
  }
]
```

Looks good, but it is not human-readable. To make it human-readable, we need to add a translator. The translator is a global variable, but don't worry, it is safe for concurrent use.

```go
func init() {
	bundle, err := errtrans.DefaultBundle()
	if err != nil {
		panic(err)
	}

	translator := errtrans.NewTranslator(errtrans.WithBundle(bundle))
	goval.SetErrorTranslator(translator)
}
```

Just execute again, and this will be the final error, with the key and a human-readable error message:

```json
[
  {
    "key":"name",
    "err":"This field is required."
  },
  {
    "key":"phone",
    "err":"This field is required."
  },
  {
    "key":"age",
    "err":"Value must be greater than 17."
  },
  {
    "key":"social_media_list",
    "err":[
      [
        {
          "key":"name",
          "err":"This field is required."
        },
        {
          "key":"link",
          "err":"This field is required."
        }
      ]
    ]
  }
]
```

### Zero Reflection
This package utilizes a new feature in Go called "Generic" to eliminate the need for the `reflect` package.

### Support for Customizable and Translatable Error Messages
As demonstrated in the previous example, we can create our own implementation of the translator by implementing the `goval.ErrorTranslator` interface. Each validation rule is already aware of this, which is why every rule requires the `context.Context` as the first argument.

We can provide the active language to the validator through the context, as shown below:

```go
ctx := context.Background()
ctx = errtrans.ContextWithLanguage(ctx, "en")
err := goval.Execute(ctx,
    goval.Named[string]("name", req.Name, goval.String().Required().Min(2).Max(20)),
    goval.Named[string]("phone", req.Phone, goval.String().Required().Match(govalregex.E164)),
    goval.Named[uint]("age", req.Age, goval.Number[uint]().Required().Min(17)),
    goval.Named[float64]("height", req.Height, goval.Number[float64]().Required()),
    goval.Named[[]SocialMedia]("social_media_list", req.SocialMediaList, goval.Slice[SocialMedia]().Required().EachFunc(SocialMediaValidator)),
)
```

## How to Contribute?

If you would like to contribute to this project, your contributions would be greatly appreciated. To contribute, 
simply fork the project and send us a pull request. Although there is no formal format for contributing at the moment, 
we would appreciate it if you could provide a good explanation with your pull request.

When you clone this repository, please make sure to run `make setup` to install the required dependencies for development
and also to set up the `pre-commit` hooks. Additionally, when you create a commit, the `pre-commit` hooks will check if the
commit follows our standard. We use Conventional Commits.

## License

Distributed under MIT License, please see license file within the code for more details.
