package main

import "fmt"

type MyInt int

type Number interface {
	~int | ~float64
}

func main() {
	intSalaries := map[string]int{
		"Arthur": 4000,
		"David":  5000,
		"John":   1000,
	}
	intTotal := sumInt(intSalaries)
	fmt.Printf("Total (int) = %d\n", intTotal)

	f64Salaries := map[string]float64{
		"Arthur": 1000.0,
		"David":  2000.0,
		"John":   3000.0,
	}
	f64Total := sumFloat64(f64Salaries)
	fmt.Printf("Total (float64) = %v\n", f64Total)

	genericTotal1 := sum(intSalaries)
	genericTotal2 := sum(f64Salaries)
	fmt.Printf("Total (generic/int) = %v\n", genericTotal1)
	fmt.Printf("Total (generic/float64) = %v\n", genericTotal2)

	myIntSalaries := map[string]MyInt{
		"Arthur": 5000,
		"David":  10000,
		"John":   5000,
	}
	myIntTotal := sum(myIntSalaries)
	fmt.Printf("Total (generic/MyInt) = %v\n", myIntTotal)

	// comparing using comparable
	fmt.Printf("10 == 10 ? answer: %v\n", compare(10, 10))
	fmt.Printf("10 == 20 ? answer: %v\n", compare(10, 20))
}

func sumInt(m map[string]int) int {
	total := 0

	for _, value := range m {
		total += value
	}

	return total
}

func sumFloat64(m map[string]float64) float64 {
	total := 0.0

	for _, value := range m {
		total += value
	}

	return total
}

// func sum[T int | float64](m map[string]T) T { // without using the constraint Number
func sum[T Number](m map[string]T) T {
	var total T

	for _, value := range m {
		total += value
	}

	return total
}

func compare[T comparable](a, b T) bool {
	if a == b {
		return true
	}

	return false
}
