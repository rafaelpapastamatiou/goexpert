package main

import "fmt"

func main() {
	ch := make(chan string)

	go send("Hello", ch)

	receive(ch)
}

// Set channel direction to send-only
func send(msg string, ch chan<- string) {
	ch <- msg
}

// Set channel direction to receive-only
func receive(ch <-chan string) {
	fmt.Println("Received: ", <-ch)
}
