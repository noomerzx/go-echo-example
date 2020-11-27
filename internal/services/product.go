package services

import (
	"go-api/internal/models"
	"go-api/internal/repositories"
	"go-api/internal/utils"
)

// ProductService interface
type ProductService interface {
	GetAllByMerchantID(MerchantID int) ([]models.Product, error)
	CreateProduct(MerchantID int, product models.Product) (models.Product, error)
}

// ProductServ struct
type ProductServ struct {
	Repo repositories.ProductRepository
}

// NewProductService func
func NewProductService(repo repositories.ProductRepository) *ProductServ {
	return &ProductServ{Repo: repo}
}

// CreateProduct func
func (s ProductServ) CreateProduct(MerchantID int, product models.Product) (models.Product, error) {
	result, _ := s.Repo.FindAllByMerchantID(MerchantID)
	if len(result) >= 5 {
		return models.Product{}, utils.NewConflictError()
	}
	product.MerchantID = MerchantID
	return s.Repo.CreateProduct(product)
}

// GetAllByMerchantID func
func (s ProductServ) GetAllByMerchantID(MerchantID int) ([]models.Product, error) {
	return s.Repo.FindAllByMerchantID(MerchantID)
}
