package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}

	json := bytes.NewBuffer([]byte(`{"name": "Rafael"}`))

	res, err := c.Post("http://google.com", "application/json", json)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)
}
