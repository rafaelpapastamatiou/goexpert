package main

import "fmt"

func main() {
	salaries := map[string]int{
		"John": 5000,
	}

	fmt.Println("John's salary:", salaries["John"])

	delete(salaries, "John")

	fmt.Println("John's salary:", salaries["John"])

	salaries["David"] = 8000

	fmt.Println("David's salary:", salaries["David"])

	// salaries2 := make(map[string]int)
	// salaries3 := map[string]int{}

	// fmt.Println(salaries2, salaries3)

	salaries["Raphael"] = 12000
	salaries["Arthur"] = 15000

	for name, salary := range salaries {
		fmt.Printf("%s's salary is %d\n", name, salary)
	}
}
