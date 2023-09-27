package main

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/pkg-id/goval"
	"github.com/pkg-id/goval/govalregex"
)

func main() {
	test_email()
	test_numeric()
	test_rgb()
}

func test_rgb() {
	fmt.Println("\n\ntest_numeric()")

	rgbValues := []string{
		"rgb(50%, 25%, 75%))",  // invalid double ))
		"rgb(100%, 101%, 50%)", // should invalid value : 101%
		"rgb(50%, 25%, 75%)",   // valid
	}
	validator := goval.String().Required().Min(2).Match(govalregex.RGB)
	ctx := context.Background()
	for _, v := range rgbValues {
		validateErr := validator.Validate(ctx, v)
		if validateErr == nil {
			fmt.Println("govalregex RGB is valid : ", v)
		} else {
			fmt.Println("govalregex RGB isn't valid : ", v)
		}
	}
}

func test_numeric() {
	fmt.Println("\n\ntest_numeric()")

	zero := "0"     // result : invalid
	minZero := "-0" // result : valid
	validator := goval.String().Required().Min(2).Match(govalregex.Numeric)
	ctx := context.Background()
	for _, v := range []string{zero, minZero} {
		validateErr := validator.Validate(ctx, v)
		if validateErr == nil {
			fmt.Println("govalregex numeric, v is valid : ", v)
		} else {
			fmt.Println("govalregex numeric, v isn't valid : ", v)
		}
	}
}

func test_email() {
	fmt.Println("test_email()")

	email := "user@example.com." // invalid email
	validator := goval.String().Required().Min(2).Match(govalregex.Email)
	ctx := context.Background()
	validateErr := validator.Validate(ctx, email)
	if validateErr == nil {
		fmt.Println("govalregex email is valid : ", email)
	} else {
		fmt.Println("govalregex email isn't valid : ", email)
	}

	// easy way to validate email
	// or another way is use validate/v10 package
	e, parseErr := mail.ParseAddress(email)
	if parseErr == nil {
		fmt.Println("net/mail email is valid : ", e)
	} else {
		fmt.Println("net/mail email isn't valid : ", e)
	}
}
