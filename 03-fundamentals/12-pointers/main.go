package main

func main() {
	a := 10

	println(a)

	println("a's memory address:", &a)

	pointer := &a

	update(pointer, 20)

	println(a)

	println("a's memory address:", &a)
}

func update(pointer *int, newValue int) {
	*pointer = newValue
}
