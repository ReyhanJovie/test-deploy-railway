package product

import (
	"errors"
)

type Repository interface {
	GetListProduct() (products []ProductEntity, err error)
}

var (
	ErrorProductNotFound = errors.New("product not found")
)

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) ListProduct() (products []ProductEntity, err error) {
	products, err = s.repo.GetListProduct()
	if err != nil {
		return
	}
	return
}
