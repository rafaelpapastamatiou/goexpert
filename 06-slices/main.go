package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	// print slice values
	fmt.Printf("Length = %d, Capacity = %d, Slice = %v\n\n", len(s), cap(s), s)

	// until index 0
	fmt.Printf("Length = %d, Capacity = %d, Slice = %v\n\n", len(s[:0]), cap(s[:0]), s[:0])

	// from index 2 until the end
	fmt.Printf("Length = %d, Capacity = %d, Slice = %v\n\n", len(s[2:]), cap(s[2:]), s[2:])

	// from index 2 until index 3
	fmt.Printf("Length = %d, Capacity = %d, Slice = %v\n\n", len(s[2:4]), cap(s[2:4]), s[2:4])

	// appending new values
	s = append(s, 6, 7, 8, 9, 10)

	fmt.Printf("Length = %d, Capacity = %d, Slice = %v\n", len(s), cap(s), s)
}
