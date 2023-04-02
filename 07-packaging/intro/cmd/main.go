package main

import (
	"fmt"

	"github.com/rafaelpapastamatiou/goexpert/07-packaging/intro/math"
)

func main() {
	m := math.Math{A: 1, B: 3 /* c: 1 */} // ! C IS NOT EXPORTED FROM math
	fmt.Println(m.Add())
	// fmt.Println("Hello World")

	// ! math2 is not exported
	// m2 := math.math2{A: 5, B: 5}
	m2 := math.NewMath2(5, 5)
	fmt.Printf("%+v\n", m2)
	// ! math2.c is not exported
	// println(m2.c)
}
