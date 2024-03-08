package main

// Main thread (01)
func main() {
	channel := make(chan string) // Create an empty channel

	// Thread 02
	go func() {
		channel <- "Hello World!" // Send a message to the channel (channel is full now)
	}()

	msg := <-channel //  Receive a message from the channel (channel is empty now)

	println(msg)
}
