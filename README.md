# Goval
[![GoDoc](https://godoc.org/github.com/pkg-id/env?status.svg)](https://godoc.org/github.com/pkg-id/goval)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/pkg-id/goval/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/pkg-id/goval)](https://goreportcard.com/report/github.com/pkg-id/goval)
[![codecov](https://codecov.io/gh/pkg-id/goval/branch/develop/graph/badge.svg?token=FCC0ZNH1D7)](https://codecov.io/gh/pkg-id/goval)

**Goval** or **Go Validator** is a package for value validation in Go.

The objective of this package is to provide a simple and easy-to-use validation library for Go that is easy to integrate
into existing projects, and easy for developers to use.
The package should be easy to extend and should provide a simple way to add custom validators, with no magic happening
behind the scenes to make it easy to understand and debug as well.

This package is designed to enhance the capabilities of the Go function as a first-class citizen. That means everything
in this package is built using function composition. Each validation rule is a simple function that can be chained (
composed) together to create complex validation logic. This package is also designed to avoid using reflection as much
as possible and is safe for concurrent use.

## How to Contribute?

The package is still under development and requires more validation rules to be implemented. If you would like to
contribute to this project, your contributions would be greatly appreciated. To contribute, simply fork the project and
send us a pull request. Although there is no formal format for contributing at the moment, we would appreciate it if you
could provide a good explanation with your pull request.

When you clone this repository, please make sure to run `make setup` to install the required dependencies for development
and also to set up the `pre-commit` hooks. Additionally, when you create a commit, the `pre-commit` hooks will check if the
commit follows our standard. We use Conventional Commits.

## Feature

- Composable Validation Rule
- Lightweight and easy to integrate with any Go application
- Support Generics
- No Reflection
- Concurrent Safe

## Example

```go
package main

import (
	"context"
	"fmt"
	"github.com/pkg-id/goval"
	"github.com/pkg-id/goval/govalregex"
)

type Product struct {
	ID            int64             `json:"id"`
	Price         float64           `json:"price"`
	Quantity      uint              `json:"quantity"`
	Note          *string           `json:"note"`
	Customization map[string]string `json:"customization"`
	OptionIndexes []int             `json:"option_indexes"`
}

func ProductValidator(p Product) goval.Validator {
	return goval.ValidatorFunc(func(ctx context.Context) error {
		return goval.Execute(ctx,
			goval.Named("id", p.ID, goval.Number[int64]().Required()),
			goval.Named("price", p.Price, goval.Number[float64]().Required()),
			goval.Named("quantity", p.Quantity, goval.Number[uint]().Required().Min(1).Max(10)),
			goval.Named("note", p.Note, goval.Ptr[string]().Optional(goval.String().Required())),
			goval.Named("customization", p.Customization, goval.Map[string, string]().Required().Each(goval.String().Required())),
			goval.Named("option_indexes", p.OptionIndexes, goval.Slice[int]().Required().Each(goval.Number[int]().Required().Min(0).Max(5))),
		)
	})
}

type Order struct {
	ID       int64     `json:"id"`
	UserID   int64     `json:"user_id"`
	Products []Product `json:"products"`
	Coupon   *string   `json:"coupon,omitempty"`
}

func main() {
	var order Order
	order.Products = []Product{{}}

	ctx := context.Background()
	err := goval.Execute(ctx,
		goval.Named("id", order.ID, goval.Number[int64]().Required()),
		goval.Named("user_id", order.ID, goval.Number[int64]().Required()),
		goval.Named("coupon", order.Coupon, goval.Ptr[string]().Optional(goval.String().Required().Match(govalregex.AlphaNumeric))),
		goval.Named("products", order.Products, goval.Each[Product](ProductValidator)),
	)
	fmt.Println(err)
}
```

## License

Distributed under MIT License, please see license file within the code for more details.
