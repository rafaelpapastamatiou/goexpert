package main

import (
	"fmt"

	"github.com/rafaelpapastamatiou/goexpert/09-apis/configs"
)

func main() {
	configs.LoadConfig("../../")
	cfg := configs.Config()
	fmt.Printf("%+v", *cfg)
}
