package main

import (
	"fmt"
	"goexpert-18/mymath"

	"github.com/google/uuid"
)

func main() {
	total := mymath.Sum(10, 20)

	// can´t use because mymath.div is not exported
	// to export something, the first letter must be capitalized

	// result := mymath.div(50, 5)

	fmt.Printf("10 + 20 = %v\n", total)

	symbols := mymath.NewOperationSymbols()

	println(symbols.Sum)
	// println(symbols.sub) // sub is not exported
	println(symbols.Div)
	// println(symbols.mult) // mult is not exported

	id := uuid.New()
	println(id.String())
}
