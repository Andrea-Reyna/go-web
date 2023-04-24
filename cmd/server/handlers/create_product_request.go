package handlers

import "github.com/Andrea-Reyna/go-web/internal/domain"

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	CodeValue   string  `json:"code_value" binding:"required"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

func (request CreateProductRequest) ToDomain() domain.Product {
	return domain.Product{
		Name:        request.Name,
		Quantity:    request.Quantity,
		CodeValue:   request.CodeValue,
		IsPublished: request.IsPublished,
		Expiration:  request.Expiration,
		Price:       request.Price,
	}
}
