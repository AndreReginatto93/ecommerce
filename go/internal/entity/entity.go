package entity

import "github.com/google/uuid"

type Category struct {
	Id   string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		Id:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	Id          string
	Name        string
	Description string
	Price       float64
	CategoryId  string
	ImageUrl    string
}

func NewProduct(name, description string, price float64, categoryId, imageUrl string) *Product {
	return &Product{
		Id:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryId:  categoryId,
		ImageUrl:    imageUrl,
	}
}
