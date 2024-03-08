package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d - Task '%s' is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// start 3 tasks (go routines/green threads)
	go task("a")
	go task("b")
	go func() {
		fmt.Println("Task 'c' is running")
	}()

	// wait forever to prevent the main go routine from exiting
	// and killing the other tasks/go routines
	for true {
	}
}
