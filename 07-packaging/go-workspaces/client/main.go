package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/rafaelpapastamatiou/goexpert/07-packaging/go-workspaces/math"
)

func main() {
	m := math.NewMath2(5, 5)
	fmt.Printf("%+v", *m)

	id := uuid.New().String()
	println(id)
}
