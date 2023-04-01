package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product // ! HAS MANY
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int // ! BELONGS TO
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductId int
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

	// ! CREATE SERIALNUMBER
	serialNumber := SerialNumber{
		Number:    "123456",
		ProductId: 1,
	}
	db.Create(&serialNumber)

	// ! PRELOAD
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	// ! PRINT CATEGORIES
	for _, c := range categories {
		fmt.Printf("Category: %+v", c)
	}
}
