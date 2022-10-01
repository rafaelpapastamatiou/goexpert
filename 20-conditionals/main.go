package main

func main() {
	a := 1
	b := 2

	if a > b && a != 0 {
		println(a)
	} else {
		println(b)
	}

	switch a {
	case 1:
		{
			println("a")
		}

	case 2:
		println("b")

	default:
		println("default")
	}
}
