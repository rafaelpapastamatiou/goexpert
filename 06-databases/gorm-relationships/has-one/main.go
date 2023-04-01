package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber // ! HAS ONE
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int // ! BELONGS TO
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert06?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// ! CREATE CATEGORY
	category := Category{
		Name: "Hardware",
	}
	db.Create(&category)

	// ! CREATE PRODUCT
	product := Product{
		Name:       "RTX 4090",
		Price:      11000.00,
		CategoryID: category.ID,
	}
	db.Create(&product)

	// ! CREATE SERIAL NUMBER
	serialNumber := SerialNumber{
		Number:    "123456",
		ProductID: product.ID,
	}
	db.Create(&serialNumber)

	// ! PRELOAD
	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}
}
