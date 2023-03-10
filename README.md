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
	"github.com/pkg-id/goval"
)

func main(){
	err := goval.String().
		Required().
		WithValue("a").
		Validate(context.Background())
	if err != nil {
		println(err)
	}
}

```

## Contributing

---

## License

---
Distributed under MIT License, please see license file within the code for more details.