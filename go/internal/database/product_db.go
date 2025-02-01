package database

import (
	"database/sql"

	"github.com/andrereginatto93/ecommerce/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProduct(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (cd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := cd.db.Query("select id, name, price, category_id, image_url from products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.ImageUrl); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (cd *ProductDB) CreateProduct(product *entity.Product) (string, error) {
	_, err := cd.db.Exec("INSERT INTO products (id, name, price, category_id, image_url) VALUES (?, ?, ?, ?, ?)",
		product.Id, product.Name, product.Price, product.CategoryId, product.ImageUrl)
	if err != nil {
		return "", err
	}

	return product.Id, nil
}

func (cd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	err := cd.db.QueryRow("Select id, name, price, category_id, image_url from products where id = ?", id).
		Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.ImageUrl)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (cd *ProductDB) GetProductByCategoryId(categoryId string) ([]*entity.Product, error) {
	rows, err := cd.db.Query("Select id, name, price, category_id, image_url from products where category_id = ?", categoryId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.ImageUrl); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
