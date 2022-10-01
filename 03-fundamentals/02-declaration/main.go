package main

const a = "Hello, World!"

var b bool
var c int
var d string
var e float64

func main() {
	// const
	println("a (const): ", a)

	// bool
	println("b (bool): ", b)
	b = true
	println("b (bool):", b)

	// int
	println("c (int):", c)
	c = 20
	println("c (int):", c)

	// string
	println("d (string):", d)
	d = "String"
	println("d (string):", d)

	// float64
	println("e (float64):", e)
	e = 35.00
	println("e (float64):", e)

	// short-hand syntax
	f := "short-hand string"
	println("f (short-hand / string):", f)
}
