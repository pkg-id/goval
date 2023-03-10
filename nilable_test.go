package goval_test

import (
	"context"
	"fmt"
	"github.com/pkg-id/goval"
)

func ExampleNil() {
	var s *string
	var n *int
	var f *float64

	ctx := context.Background()
	err := goval.Nil[string]().Required().Build(s).Validate(ctx)
	fmt.Println("Nil[string]:", err) // Nil[string]: is required

	err = goval.Nil[int]().Required().Build(n).Validate(ctx)
	fmt.Println("Nil[int]:", err) // Nil[int]: is required

	err = goval.Nil[float64]().Required().Build(f).Validate(ctx)
	fmt.Println("Nil[float64]:", err) // Nil[float64]: is required

	err = goval.Nil[float64]().Optional(goval.Number[float64]().Required().Max(2)).Build(f).Validate(ctx)
	fmt.Println("Nil[float64].Optional", err == nil)

	s = new(string)
	n = new(int)
	f = new(float64)

	*s = "a"
	*n = 1
	*f = 3.14

	err = goval.Nil[string]().Required().Next(goval.String().Required().Min(2)).Build(s).Validate(ctx)
	fmt.Println("Nil[string].Next:", err) // Nil[string].Next: length must be at least 2 characters

	err = goval.Nil[int]().
		Required().
		Next(goval.Number[int]().Required().Min(2)).
		Build(n).
		Validate(ctx)
	fmt.Println("Nil[int].Next:", err) // Nil[int].Next: must be greater than 2

	err = goval.Nil[float64]().Optional(goval.Number[float64]().Required().Max(2)).Build(f).Validate(ctx)
	fmt.Println("Nil[float64].Optional:", err) // Nil[float64].Next: must be less than 2

	// Output:
	// Nil[string]: is required
	// Nil[int]: is required
	// Nil[float64]: is required
	// Nil[float64].Optional true
	// Nil[string].Next: length must be at least 2 characters
	// Nil[int].Next: must be greater than 2
	// Nil[float64].Optional: must be less than 2
}
