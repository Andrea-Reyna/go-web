package products

import (
	"errors"

	"github.com/Andrea-Reyna/go-web/internal/domain"
)

var (
	ErrProductAlreadyExists = errors.New("product already exist")
	ErrProductNotFound      = errors.New("product not found")
)

type Repository interface {
	Create(product *domain.Product) error
	GetAll() ([]domain.Product, error)
	FindById(id int) (domain.Product, error)
	Search(priceGt float64) ([]domain.Product, error)
	Update(product *domain.Product) error
	UpdateName(id int, name string) (domain.Product, error)
	Delete(id int) error
	ConsumerPrice(list []int) ([]domain.Product, error)
}
