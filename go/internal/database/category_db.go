package database

import (
	"database/sql"

	"github.com/andrereginatto93/ecommerce/goapi/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (cd *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := cd.db.Query("select id, name from categories")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category

	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.Id, category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := cd.db.Exec("INSERT INTO categories (id, name) VALUES (?, ?)", category.Id, category.Name)
	if err != nil {
		return "", err
	}

	return category.Id, nil
}

func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var category entity.Category
	err := cd.db.QueryRow("Select id, name from categories where id = ?", id).Scan(&category.Id, &category.Name)

	if err != nil {
		return nil, err
	}

	return &category, nil
}
