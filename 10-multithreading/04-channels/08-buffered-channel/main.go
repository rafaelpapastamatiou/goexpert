package main

func main() {
	// Create a buffered channel with a capacity of 5
	ch := make(chan string, 5)

	ch <- "message 1"
	ch <- "message 2"
	ch <- "message 3"
	ch <- "message 4"
	ch <- "message 5"

	println(<-ch)
	println(<-ch)
	println(<-ch)
	println(<-ch)
	println(<-ch)
}
