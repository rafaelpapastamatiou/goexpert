package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/config"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/infra/database"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/infra/http/handler"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig("../../")
	//cfg := configs.Config()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productsRepository := database.NewGormProductsRepository(db)

	newProductHandler := handler.NewProductHandler(productsRepository)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/products", newProductHandler.GetProducts)
	r.Get("/products/{id}", newProductHandler.GetProduct)
	r.Post("/products", newProductHandler.CreateProduct)
	r.Put("/products/{id}", newProductHandler.UpdateProduct)
	r.Delete("/products/{id}", newProductHandler.DeleteProduct)

	http.ListenAndServe(":8000", r)
}
