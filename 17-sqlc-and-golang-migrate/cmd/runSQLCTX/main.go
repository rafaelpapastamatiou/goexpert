package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/rafaelpapastamatiou/goexpert/17-sqlc/internal/db"
)

type BaseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func (bdb *BaseDB) CallTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := bdb.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	qtx := db.New(tx)

	err = fn(qtx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

type CourseDB struct {
	BaseDB
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		BaseDB: BaseDB{
			dbConn:  dbConn,
			Queries: db.New(dbConn),
		},
	}
}

type CreateCourseParams struct {
	Name         string
	Description  string
	CategoryName string
}

func (cdb *CourseDB) CreateCourse(ctx context.Context, data CreateCourseParams) error {
	return cdb.CallTx(ctx, func(q *db.Queries) error {
		// Check if category exists
		category, err := q.GetCategoryByName(ctx, data.CategoryName)

		if err != nil && err != sql.ErrNoRows {
			return err
		}

		if err == sql.ErrNoRows {
			// Create category
			err = q.CreateCategory(ctx, db.CreateCategoryParams{
				ID:          uuid.New().String(),
				Name:        data.CategoryName,
				Description: sql.NullString{String: data.CategoryName, Valid: true},
			})
			if err != nil {
				return err
			}

			// Get created category
			category, err = q.GetCategoryByName(ctx, data.CategoryName)
			if err != nil {
				return err
			}
		}

		// Create course
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          uuid.New().String(),
			Name:        data.Name,
			Description: sql.NullString{String: data.Description, Valid: true},
			CategoryID:  category.ID,
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert-17-sqlc")
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	// courseDB := NewCourseDB(dbConn)

	// createCourseData := CreateCourseParams{
	// 	Name:         "Course TX",
	// 	Description:  "Course TX Description",
	// 	CategoryName: "Category TX",
	// }

	// err = courseDB.CreateCourse(ctx, createCourseData)
	// if err != nil {
	// 	println("Error creating course using transaction")
	// 	panic(err)
	// }

	queries := db.New(dbConn)
	courses, err := queries.ListCourses(ctx)
	if err != nil {
		println("Error listing courses")
		panic(err)
	}

	println("Courses:")
	for _, course := range courses {
		println("ID:", course.ID)
		println("Name:", course.Name)
		println("Description:", course.Description.String)
		println("Category ID:", course.CategoryID)
		println("Category Name:", course.CategoryName)
		println("Category Description:", course.CategoryDescription.String)
		println()
	}
}
