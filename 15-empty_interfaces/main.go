package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)

	a := 20
	b := map[string]int{
		"i": 20,
		"j": 30,
	}

	showType(a)
	showType(b)
}

func showType(t interface{}) {
	fmt.Printf("Var's type is %T and the value is %v\n", t, t)
}
