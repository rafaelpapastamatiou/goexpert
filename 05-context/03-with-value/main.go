package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "Rafael")

	token := ctx.Value("name")
	fmt.Println(token)
}
