package main

import "fmt"

func main() {
	totalMultipliedBy2 := func(a, b int) int {
		return sum(a, b) * 2
	}(5, 5)

	fmt.Println(totalMultipliedBy2)
}

func sum(a, b int) int {
	return a + b
}
