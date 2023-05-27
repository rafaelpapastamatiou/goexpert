package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/config"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/infra/database"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/infra/http/handler"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config.LoadConfig("../../")
	cfg := config.Config()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productsRepository := database.NewGormProductsRepository(db)
	usersRepository := database.NewGormUsersRepository(db)

	productHandler := handler.NewProductHandler(productsRepository)
	userHandler := handler.NewUserHandler(
		usersRepository,
		cfg.AuthToken,
		cfg.JWTExpiresIn,
	)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.AuthToken))
		r.Use(jwtauth.Authenticator)

		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Post("/", productHandler.CreateProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/auth", userHandler.AuthenticateUser)

	http.ListenAndServe(":8000", r)
}
