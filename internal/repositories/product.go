package repositories

import (
	"go-api/internal/app"
	"go-api/internal/models"
)

// ProductRepository interface
type ProductRepository interface {
	FindAllByMerchantID(MerchantID int) ([]models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
}

// ProductRepo struct
type ProductRepo struct {
	cf *app.Config
}

// NewProductRepo func
func NewProductRepo(cf *app.Config) *ProductRepo {
	return &ProductRepo{cf: cf}
}

// CreateProduct func
func (r ProductRepo) CreateProduct(product models.Product) (models.Product, error) {
	if len(r.cf.Db.Products) != 0 {
		lastProduct := r.cf.Db.Products[len(r.cf.Db.Products)-1]
		product.ID = lastProduct.ID + 1
	} else {
		product.ID = 1
	}
	r.cf.Db.Products = append(r.cf.Db.Products, product)
	return product, nil
}

// FindAllByMerchantID func
func (r ProductRepo) FindAllByMerchantID(MerchantID int) ([]models.Product, error) {
	var products []models.Product
	for _, p := range r.cf.Db.Products {
		if p.MerchantID == MerchantID {
			products = append(products, p)
		}
	}
	return products, nil
}
