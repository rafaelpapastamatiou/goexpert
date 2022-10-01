package main

func main() {
	for a := 0; a < 10; a++ {
		println("a", a)
	}

	numbers := []string{"one", "two", "three", "four", "five"}

	for key, value := range numbers {
		println(key, ":", value)
	}

	for _, value := range numbers {
		println(value)
	}

	for key := range numbers {
		println(key)
	}

	b := 0

	for b < 10 {
		println(b)
		b++
	}

	c := 0

	for {
		println(c)

		c++

		if c > 10 {
			break
		}
	}
}
