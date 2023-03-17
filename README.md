# Goval

Goval is a powerful and easy-to-use library that provides a set of functions for validating
various types of data in Go programming language.
This library is designed to simplify the process of data validation and reduce the likelihood of errors and vulnerabilities in applications.

## Design

The Engineering Design behind this package can be accessed in `DESIGN.md`.

## Features

---

- Validate strings, numbers, dates, structs, and more
- Customizable error messages and validation rules
- Lightweight and easy to integrate with any Go application
- Built-in support for validating structs

## Install

---

Use go get
```shell
go get github.com/pkg-id/goval
```

Then, import to your own code
```go
import "github.com/pkg-id/goval"
```

## Usage

---

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

## Contributing

---

## License

---
Distributed under MIT License, please see license file within the code for more details.
