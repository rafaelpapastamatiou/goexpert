package main

import (
	"fmt"
)

func main() {
	fmt.Println("First line")

	defer fmt.Println("Last line")  // runs before function return
	defer fmt.Println("Third line") // runs before the previous defer (LIFO)

	fmt.Println("Second line")
}
