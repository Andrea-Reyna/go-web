package products

import (
	"time"

	"github.com/Andrea-Reyna/go-web/internal/domain"
)

type DefaultService struct {
	Storage Repository
}

func (service DefaultService) Create(product *domain.Product) error {
	err := service.validations(product)
	if err != nil {
		return err
	}
	err = service.Storage.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (service DefaultService) Update(product *domain.Product) error {
	err := service.validations(product)
	if err != nil {
		return err
	}
	err = service.Storage.Update(product)
	if err != nil {
		return err
	}
	return nil
}

func (service DefaultService) UpdateName(id int, name string) (domain.Product, error) {
	if name == "" {
		return domain.Product{}, ErrInvalidData
	}

	newProduct, err := service.Storage.UpdateName(id, name)
	if err != nil {
		return domain.Product{}, err
	}
	return newProduct, err
}

func (service DefaultService) GetAll() ([]domain.Product, error) {
	products, err := service.Storage.GetAll()
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}

func (service DefaultService) FindById(id int) (domain.Product, error) {
	products, err := service.Storage.FindById(id)
	if err != nil {
		return domain.Product{}, ErrProductNotFound

	}
	return products, nil
}

func (service DefaultService) Search(priceGt float64) ([]domain.Product, error) {
	if priceGt <= 0 {
		return []domain.Product{}, ErrInvalidData
	}
	products, err := service.Storage.Search(priceGt)
	if err != nil {
		return []domain.Product{}, ErrInternalServerError
	}
	return products, nil
}

func (service DefaultService) Delete(id int) error {
	err := service.Storage.Delete(id)
	if err != nil {
		return ErrProductNotFound
	}
	return nil
}

func (service DefaultService) ConsumerPrice(list []int) (domain.ProductsConsumer, error) {
	filterProducts, err := service.Storage.ConsumerPrice(list)
	if err != nil {
		return domain.ProductsConsumer{}, ErrInternalServerError
	}
	var productsConsumer domain.ProductsConsumer
	productsConsumer.Products = filterProducts

	for _, value := range filterProducts {
		productsConsumer.TotalPrice += value.Price
	}

	if len(filterProducts) < 10 {
		productsConsumer.TotalPrice = productsConsumer.TotalPrice * 1.21
	}

	if len(filterProducts) > 10 && len(filterProducts) < 20 {
		productsConsumer.TotalPrice = productsConsumer.TotalPrice * 1.17
	}
	if len(filterProducts) > 20 {
		productsConsumer.TotalPrice = productsConsumer.TotalPrice * 1.15
	}

	return productsConsumer, nil
}

func (service DefaultService) validations(product *domain.Product) error {
	products, err := service.Storage.GetAll()
	if err != nil {
		return err
	}

	for _, prod := range products {
		if product.CodeValue == prod.CodeValue && product.ID != prod.ID {
			return ErrProductAlreadyExists
		}
	}

	_, err = time.Parse("02/01/2006", product.Expiration)
	if err != nil {
		return ErrFormateDate
	}

	return nil
}
