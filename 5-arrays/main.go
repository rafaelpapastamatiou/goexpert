package main

func main() {
	var array [5]int

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	println("first element", array[0])
	println("last element", array[len(array)-1])

	for i, v := range array {
		println("Index:", i, "Value:", v)
	}
}
