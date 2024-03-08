package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d - Task '%s' is running\n", i, name)
		time.Sleep(1 * time.Second)
	}

	// indicate that the task is done
	wg.Done()
}

func main() {
	// create a wait group
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)

	// start 3 tasks (go routines/green threads)
	go task("a", &waitGroup)
	go task("b", &waitGroup)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Task 'c' is running")

		// indicate that the task is done
		wg.Done()
	}(&waitGroup)

	// wait until all tasks are done to exit the program
	waitGroup.Wait()
}
