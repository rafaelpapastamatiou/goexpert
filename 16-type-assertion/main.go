package main

func main() {
	var x interface{} = "Hello, World!"

	println(x)
	println(x.(string))

	res, ok := x.(int)

	if ok == false {
		println("x can´t be converted to int")
	} else {
		println(res)
	}

	res = x.(int)

	println(res)
}
