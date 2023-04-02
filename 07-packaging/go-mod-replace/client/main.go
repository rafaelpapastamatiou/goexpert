package main

import (
	"fmt"

	"github.com/rafaelpapastamatiou/goexpert/07-packaging/go-mod-replace/math"
)

func main() {
	m := math.NewMath2(5, 5)
	fmt.Printf("%+v", m)
}
