package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}

	productUsecase := NewProductUseCase(db)

	product, err := productUsecase.GetProductByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product ID: %d, Name: %s\n", product.ID, product.Name)
}
