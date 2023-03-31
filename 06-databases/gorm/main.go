package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert06?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// ! CREATE
	// db.Create(&Product{
	// 	Name:  "Headphone",
	// 	Price: 1500,
	// })

	// ! BULK CREATE
	// products := []Product{
	// 	{Name: "Keyboard", Price: 800},
	// 	{Name: "Notebook", Price: 6000},
	// 	{Name: "Mouse", Price: 800},
	// }
	// db.Create(&products)

	// ! FIND ONE
	// var product Product
	// db.First(&product, 1)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// ! FIND
	// var products []Product
	// db.Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// ! FIND WITH LIMIT
	// var products []Product
	// db.Limit(2).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// ! FIND WITH LIMIT AND OFFSET
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// ! WHERE
	// var products []Product
	// db.Where("price > ?", 1000).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// ! WHERE + LIKE
	// var products []Product
	// db.Where("name LIKE ?", "%board%").Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// ! UPDATE
	// var p Product
	// db.First(&p, 1)
	// p.Name = "Earphone"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2)

	// ! DELETE
	// var p Product
	// db.First(&p, 1)
	// fmt.Println(p)
	// db.Delete(&p)
}
