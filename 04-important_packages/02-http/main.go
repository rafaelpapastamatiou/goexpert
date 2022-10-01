package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	response.Body.Close()
}
