-- name: ListCategories :many
SELECT * from categories;

-- name: GetCategory :one
SELECT * from categories WHERE id = ?;

-- name: GetCategoryByName :one
SELECT * from categories WHERE name = ?;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) VALUES (?, ?, ?);

-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ? WHERE id = ?;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = ?;

-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, category_id) VALUES (?, ?, ?, ?);

-- name: ListCourses :many
SELECT 
  c.*,
  cat.name as category_name,
  cat.description as category_description
FROM courses c
INNER JOIN categories cat
ON c.category_id = cat.id;