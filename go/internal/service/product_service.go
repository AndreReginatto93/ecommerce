package service

import (
	"github.com/andrereginatto93/ecommerce/goapi/internal/database"
	"github.com/andrereginatto93/ecommerce/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (cs *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := cs.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (cs *ProductService) CreateProduct(name, description string, price float64, categoryId, imageUrl string) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, categoryId, imageUrl)

	_, err := cs.ProductDB.CreateProduct(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (cd *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := cd.ProductDB.GetProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (cd *ProductService) GetProductByCategory(categoryId string) ([]*entity.Product, error) {
	products, err := cd.ProductDB.GetProductByCategoryId(categoryId)

	if err != nil {
		return nil, err
	}

	return products, nil
}
