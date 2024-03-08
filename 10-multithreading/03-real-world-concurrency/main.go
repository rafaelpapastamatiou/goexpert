package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{} // Create a mutex to lock the number variable

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock() // Lock the number variable
		// number++
		// m.Unlock() // Unlock the number variable

		// Use the atomic package to increment the number variable
		// avoiding the need to lock and unlock the variable manually
		atomic.AddUint64(&number, 1)

		w.Write([]byte(
			fmt.Sprintf("You are visitor number %d\n", number),
		))
	})

	http.ListenAndServe(":8000", nil)
}
