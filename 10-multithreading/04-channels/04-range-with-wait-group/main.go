package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	// Start a goroutine to send messages to the channel
	go publish(ch, &wg, 10)

	// Start a goroutine to process messages from the channel
	go process(ch, &wg)

	// Wait for all messages to be processed
	wg.Wait()
}

func publish(ch chan int, wg *sync.WaitGroup, amount int) {
	// Add {amount} to the wait group counter
	wg.Add(amount)

	// Send {amount} messages to the channel
	for i := 0; i < amount; i++ {
		ch <- i
	}

	// Close the channel after sending all messages
	close(ch)
}

func process(ch chan int, wg *sync.WaitGroup) {
	// Process messages from the channel until it's closed
	for i := range ch {
		fmt.Printf("Received %d\n", i)

		// Indicate that the message was processed
		wg.Done()
	}
}
