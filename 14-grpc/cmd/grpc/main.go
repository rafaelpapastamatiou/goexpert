package main

import (
	"database/sql"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rafaelpapastamatiou/goexpert/14-grpc/internal/database"
	"github.com/rafaelpapastamatiou/goexpert/14-grpc/internal/pb"
	"github.com/rafaelpapastamatiou/goexpert/14-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	categoryService := service.NewCategoryService(categoryDb)
	courseService := service.NewCourseService(courseDb)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	pb.RegisterCourseServiceServer(grpcServer, courseService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
