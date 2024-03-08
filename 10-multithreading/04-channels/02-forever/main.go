package main

import "time"

func main() {
	forever := make(chan bool) // Create an empty channel

	go func() {
		println("Waiting...")
		time.Sleep(3 * time.Second)
		forever <- true // Send a message to the channel, to avoid a dead lock
	}()

	<-forever // Channel is empty, so this will throw a dead lock error
}
