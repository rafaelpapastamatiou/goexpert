package usecase

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rafaelpapastamatiou/goexpert/18-uow-unit-of-work/internal/db"
	"github.com/rafaelpapastamatiou/goexpert/18-uow-unit-of-work/internal/repository"
	"github.com/rafaelpapastamatiou/goexpert/18-uow-unit-of-work/pkg/uow"
	"github.com/stretchr/testify/assert"
)

func TestAddCourseUsecase(t *testing.T) {
	sqlDb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert-18-uow")
	assert.NoError(t, err)
	defer sqlDb.Close()

	sqlDb.Exec("DROP TABLE IF EXISTS `courses`;")
	sqlDb.Exec("DROP TABLE IF EXISTS `categories`;")

	sqlDb.Exec("CREATE TABLE IF NOT EXISTS `categories` (`id` varchar(255) NOT NULL, `name` varchar(255) NOT NULL, `description` varchar(255) DEFAULT NULL, PRIMARY KEY (`id`));")
	sqlDb.Exec("CREATE TABLE IF NOT EXISTS `courses` (`id` varchar(255) NOT NULL, `name` varchar(255) NOT NULL, `description` varchar(255) DEFAULT NULL, `category_id` varchar(255) NOT NULL, FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`), PRIMARY KEY (`id`));")

	input := AddCourseUsecaseInput{
		CourseName:        "Test Course 1",
		CourseDescription: "Test Course Description 1",
		CategoryName:      "Test Category 1",
	}

	ctx := context.Background()

	unitOfWork := uow.NewUnitOfWork(sqlDb)

	unitOfWork.Register("category", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(sqlDb)
		repo.Queries = db.New(tx)
		return repo
	})

	unitOfWork.Register("course", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(sqlDb)
		repo.Queries = db.New(tx)
		return repo
	})

	usecase := NewAddCourseUsecase(unitOfWork)
	err = usecase.Execute(ctx, input)
	assert.NoError(t, err)
}
