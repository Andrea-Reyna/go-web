package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Andrea-Reyna/go-web/internal/domain"
)

func LoadFile() ([]domain.Product, error) {
	products := []domain.Product{}
	file, err := os.Open(os.Getenv("FILE"))
	if err != nil {
		return products, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)
	if err != nil {
		return products, fmt.Errorf("error decoding data: %w", err)
	}
	return products, nil
}

func SaveProducts(products []domain.Product) error {
	file, err := os.Open(os.Getenv("FILE"))
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	data, err := json.Marshal(products)
	if err != nil {
		return errors.New("error converting text to json")
	}

	err = os.WriteFile(os.Getenv("FILE"), data, 0644)
	if err != nil {
		return errors.New("error writting file")
	}
	return nil
}
