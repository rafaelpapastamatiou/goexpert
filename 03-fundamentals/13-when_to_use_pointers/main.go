package main

import "fmt"

func main() {
	v1 := 10
	v2 := 20

	fmt.Printf("v1=%d, v2=%d\n", v1, v2)

	r1 := sum(v1, v2)

	fmt.Printf("%d + %d = %d\n", v1, v2, r1)

	v3 := 25
	v4 := 50

	fmt.Printf("v3=%d, v4=%d\n", v3, v4)

	r2 := sumChangingInputValues(&v3, &v4)

	fmt.Printf("%d + %d = %d\n", v3, v4, r2)

}

func sum(a, b int) int {
	return a + b
}

func sumChangingInputValues(a, b *int) int {
	*a *= 2
	*b *= 2

	return *a + *b
}
