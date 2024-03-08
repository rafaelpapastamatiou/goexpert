package main

import "time"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		// Send a message to the channel instantly
		c1 <- 1
	}()

	go func() {
		// Send a message to the channel after 1 second
		time.Sleep(1 * time.Second)
		c2 <- 2
	}()

	// Select the first channel that has data available
	select {
	case msg1 := <-c1:
		println("Received from c1:", msg1)

	case msg2 := <-c2:
		println("Received from c2:", msg2)

	// Timeout after 3 seconds
	case <-time.After(3 * time.Second):
		println("Timeout")

	default:
		println("None of the channels were ready")
	}
}
