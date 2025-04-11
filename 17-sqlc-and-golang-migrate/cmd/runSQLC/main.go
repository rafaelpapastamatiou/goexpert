package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/rafaelpapastamatiou/goexpert/17-sqlc/internal/db"
)

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert-17-sqlc")
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)

	// Create a new category
	categoryId := uuid.New().String()

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          categoryId,
		Name:        "Category 1",
		Description: sql.NullString{String: "Description 1", Valid: true},
	})
	if err != nil {
		fmt.Printf("Error creating category")
		panic(err)
	}

	fmt.Printf("CreateCategory - Category created with ID: %s\n\n", categoryId)

	// Get the category
	category, err := queries.GetCategory(ctx, categoryId)
	if err != nil {
		fmt.Printf("Error getting category")
		panic(err)
	}

	fmt.Printf("GetCategory - Category ID: %s, Name: %s, Description: %s\n\n", category.ID, category.Name, category.Description.String)

	// Update the category
	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          categoryId,
		Name:        "Updated Category 1",
		Description: sql.NullString{String: "Updated Description 1", Valid: true},
	})
	if err != nil {
		fmt.Printf("Error updating category")
		panic(err)
	}

	fmt.Printf("UpdateCategory - Category updated with ID: %s\n\n", categoryId)

	// Get the updated category
	category, err = queries.GetCategory(ctx, categoryId)
	if err != nil {
		fmt.Printf("Error getting updated category")
		panic(err)
	}

	fmt.Printf("GetCategory - Updated Category ID: %s, Name: %s, Description: %s\n\n", category.ID, category.Name, category.Description.String)

	// List all categories
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		fmt.Printf("Error listing categories")
		panic(err)
	}
	println("ListCategories - Categories:")
	for _, category := range categories {
		fmt.Printf("Category ID: %v, Name: %v, Description: %s\n", category.ID, category.Name, category.Description.String)
	}
}
