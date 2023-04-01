package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"` // ! MANY TO MANY
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"` // ! MANY TO MANY
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert06?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// ! CREATE CATEGORIES
	category := Category{Name: "Hardware"}
	db.Create(&category)

	category2 := Category{Name: "Eletronics"}
	db.Create(&category2)

	// ! CREATE PRODUCT
	product := Product{
		Name:       "RTX 4090",
		Price:      11000.00,
		Categories: []Category{category, category2},
	}
	db.Create(&product)

	// ! PRELOAD
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	// ! PRINT CATEGORIES
	for _, c := range categories {
		fmt.Printf("Category: %+v", c)
	}
}
