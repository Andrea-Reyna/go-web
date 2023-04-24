package products

import (
	"errors"

	"github.com/Andrea-Reyna/go-web/internal/domain"
)

var (
	ErrProductInvalid      = errors.New("invalid product")
	ErrFormateDate         = errors.New("invalid format date")
	ErrInvalidData         = errors.New("invalid data")
	ErrInternalServerError = errors.New("internal server error")
)

type Service interface {
	Create(product *domain.Product) error
	GetAll() ([]domain.Product, error)
	FindById(id int) (domain.Product, error)
	Search(priceGt float64) ([]domain.Product, error)
	Update(product *domain.Product) error
	UpdateName(id int, name string) (domain.Product, error)
	Delete(id int) error
	ConsumerPrice(list []int) (domain.ProductsConsumer, error)
}
