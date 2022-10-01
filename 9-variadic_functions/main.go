package main

import "fmt"

func main() {
	println(fmt.Sprint(sum(1, 2, 3)))

	println(fmt.Sprint(sum(10, 5, 5)))
}

func sum(numbers ...int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}

	return total
}

func sumWithInitialValue(initialValue int, numbers ...int) int {
	total := initialValue

	for _, n := range numbers {
		total += n
	}

	return total
}
