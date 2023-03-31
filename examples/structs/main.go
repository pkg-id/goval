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

func ProductValidator() goval.RuleValidatorFunc[Product] {
	return func(ctx context.Context, p Product) error {
		return goval.Execute(ctx,
			goval.Named("id", p.ID, goval.Number[int64]().Required()),
			goval.Named("price", p.Price, goval.Number[float64]().Required()),
			goval.Named("quantity", p.Quantity, goval.Number[uint]().Required().Min(1).Max(10)),
			goval.Named("note", p.Note, goval.Ptr[string]().Optional(goval.String().Required())),
			goval.Named("customization", p.Customization, goval.Map[string, string]().Required().Each(goval.String().Required())),
			goval.Named("option_indexes", p.OptionIndexes, goval.Slice[int]().Required().Each(goval.Number[int]().Required().Min(0).Max(5))),
		)
	}
}

type Order struct {
	ID       int64     `json:"id"`
	UserID   int64     `json:"user_id"`
	Products []Product `json:"products"`
	Coupon   *string   `json:"coupon,omitempty"`
}

func init() {
	goval.SetErrorTranslator(&customTranslator{})
}

func main() {
	var order Order
	order.Products = []Product{{}}

	ctx := context.Background()
	err := goval.Execute(ctx,
		goval.Named("id", order.ID, goval.Number[int64]().Required()),
		goval.Named("user_id", order.ID, goval.Number[int64]().Required()),
		goval.Named("coupon", order.Coupon, goval.Ptr[string]().Optional(goval.String().Required().Match(govalregex.AlphaNumeric))),
		goval.Named("products", order.Products, goval.Each[Product](ProductValidator())),
	)
	fmt.Println(err)
}

type customTranslator struct {
}

func (c *customTranslator) Translate(_ context.Context, err *goval.RuleError) error {
	switch err.Code {
	case goval.StringRequired:
		return goval.TextError("string is required")
	case goval.NumberRequired:
		return goval.TextError("number is required")
	default:
		return err
	}
}
