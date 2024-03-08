package main

import "fmt"

func main() {
	ch := make(chan int)

	go publish(ch, 10)

	// Read messages from the channel until it's closed
	for i := range ch {
		fmt.Printf("Received %d\n", i)
	}
}

func publish(ch chan int, amount int) {
	// Send {amount} messages to the channel
	for i := 0; i < amount; i++ {
		ch <- i
	}

	// Close the channel after sending all messages
	close(ch)
}
