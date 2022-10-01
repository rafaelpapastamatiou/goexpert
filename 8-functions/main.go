package main

import (
	"errors"
	"fmt"
)

func main() {
	c := sum(1, 2)

	log(fmt.Sprint(c))

	result1, hasRest1 := div(10, 5)
	println("10 / 5 = " + fmt.Sprint(result1, hasRest1))

	result2, hasRest2 := div(10, 3)
	println("10 / 3 = " + fmt.Sprint(result2, hasRest2))

	result3, err := multiplyBy2IfEven(2)

	if err != nil {
		println(err.Error())
	} else {
		println(fmt.Sprintf("2 * 2 = %d", result3))
	}

	result4, err := multiplyBy2IfEven(3)

	if err != nil {
		println(err.Error())
	} else {
		println(fmt.Sprintf("3 * 2 = %d", result4))
	}
}

func sum(a, b int) int {
	return a + b
}

func div(a, b int) (result int, hasRest bool) {
	r := a / b
	rest := a % b

	if rest != 0 {
		return r, true
	}

	return r, false
}

func multiplyBy2IfEven(number int) (result int, err error) {
	if number%2 != 0 {
		return 0, errors.New(fmt.Sprintf("%d is not even", number))
	}

	return number / 2, nil
}

func log(message string) {
	println(message)
}
