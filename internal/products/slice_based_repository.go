package products

import (
	"fmt"

	"github.com/Andrea-Reyna/go-web/internal/domain"
	"github.com/Andrea-Reyna/go-web/pkg/store"
)

type SliceBasedRepository struct {
	products []domain.Product
}

func NewSliceBasedRepository() (*SliceBasedRepository, error) {
	prod, err := store.LoadFile()

	if err != nil {
		return nil, fmt.Errorf("error data: %w", err)
	}
	repository := &SliceBasedRepository{
		products: prod,
	}
	return repository, nil
}

func (repository *SliceBasedRepository) Create(product *domain.Product) (err error) {
	product.ID = len(repository.products) + 1
	repository.products = append(repository.products, *product)
	//store.SaveProducts(repository.products)
	return
}

func (repository *SliceBasedRepository) Update(product *domain.Product) (err error) {
	var updated bool
	for i := range repository.products {
		if repository.products[i].ID == product.ID {
			repository.products[i] = *product
			updated = true
			break
		}
	}
	if !updated {
		return ErrProductNotFound
	}
	//store.SaveProducts(repository.products)
	return
}

func (repository *SliceBasedRepository) UpdateName(id int, name string) (domain.Product, error) {
	var updated bool
	var product domain.Product
	for i := range repository.products {
		if repository.products[i].ID == id {
			repository.products[i].Name = name
			updated = true
			break
		}
	}
	if !updated {
		return domain.Product{}, ErrProductNotFound
	}
	//store.SaveProducts(repository.products)
	return product, nil
}

func (repository *SliceBasedRepository) GetAll() ([]domain.Product, error) {
	return repository.products, nil
}

func (repository *SliceBasedRepository) FindById(id int) (domain.Product, error) {
	for i := range repository.products {
		if repository.products[i].ID == id {
			return repository.products[i], nil
		}
	}
	return domain.Product{}, ErrProductNotFound
}
func (repository *SliceBasedRepository) Search(priceGt float64) ([]domain.Product, error) {
	var filterProducts []domain.Product
	for i := range repository.products {
		if repository.products[i].Price > priceGt {
			filterProducts = append(filterProducts, repository.products[i])
		}
	}
	return filterProducts, nil
}

func (repository *SliceBasedRepository) Delete(id int) error {
	var deleted bool
	var index int
	for i := range repository.products {
		if repository.products[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return ErrProductNotFound
	}
	repository.products = append(repository.products[:index], repository.products[index+1:]...)
	store.SaveProducts(repository.products)
	return nil
}

func (repository *SliceBasedRepository) ConsumerPrice(list []int) ([]domain.Product, error) {
	var filterProducts []domain.Product
	for _, id := range list {
		for i := range repository.products {
			if repository.products[i].ID == id && repository.products[i].IsPublished {
				filterProducts = append(filterProducts, repository.products[i])
			}
		}
	}
	//store.SaveProducts(repository.products)
	return filterProducts, nil
}
