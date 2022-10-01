package main

import "fmt"

type ID int

var x ID = 10
var y string = "Hello, World!"

func main() {
	fmt.Printf("O tipo de x é %T\n", x)
	fmt.Printf("O tipo de y é %T\n", y)
}
