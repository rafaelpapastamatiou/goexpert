package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)

	workersAmount := 100
	messagesAmount := workersAmount * 10

	// Start workers
	for i := 1; i <= workersAmount; i++ {
		go worker(i, data)
	}

	for i := 0; i < messagesAmount; i++ {
		data <- i
	}

	time.Sleep(5 * time.Second)
}

func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
	}
}
