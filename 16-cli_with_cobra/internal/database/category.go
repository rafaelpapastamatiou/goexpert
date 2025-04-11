package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name, description string) (Category, error) {
	id := uuid.New().String()

	stmt, err := c.db.Prepare("INSERT INTO categories (id, name, description) values (?, ?, ?)")
	if err != nil {
		return Category{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id, name, description)
	if err != nil {
		return Category{}, err
	}

	return Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (c *Category) List() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []Category{}

	for rows.Next() {
		var id, name, description string

		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		categories = append(categories, Category{
			ID:          id,
			Name:        name,
			Description: description,
		})
	}

	return categories, nil
}

func (c *Category) FindById(id string) (Category, error) {
	stmt, err := c.db.Prepare("SELECT id, name, description FROM categories WHERE id = ?")
	if err != nil {
		return Category{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var category Category

	if err := row.Scan(&category.ID, &category.Name, &category.Description); err != nil {
		return Category{}, err
	}

	return category, nil
}
