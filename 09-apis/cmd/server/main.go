package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/config"
	_ "github.com/rafaelpapastamatiou/goexpert/09-apis/docs"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/domain/entity"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/infra/database"
	"github.com/rafaelpapastamatiou/goexpert/09-apis/internal/infra/http/handler"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert Example API
// @version         1.0
// @description     Products API with authentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Rafael Papastamatiou
// @contact.url    http://papastamatiou.com
// @contact.email  rafael@papastamatiou.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name Authorization
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
	userHandler := handler.NewUserHandler(usersRepository)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", cfg.AuthToken))
	r.Use(middleware.WithValue("jwtExpiresIn", cfg.JWTExpiresIn))
	//r.Use(middleware.Logger)
	r.Use(LogRequest)

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

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/docs/doc.json"),
	))

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
