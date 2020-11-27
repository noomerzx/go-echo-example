package services

import (
	"go-api/internal/models"
	"testing"
)

var fakeProduct = []models.Product{
	{
		ID:         1,
		Name:       "iPhone 12",
		MerchantID: 1,
		Amount:     31000.00,
		Stock:      5,
	},
	{
		ID:         2,
		Name:       "iPhone 12 mini",
		MerchantID: 1,
		Amount:     25000.00,
		Stock:      15,
	},
	{
		ID:         3,
		Name:       "iPhone 12 Pro Max",
		MerchantID: 2,
		Amount:     51000.00,
		Stock:      2,
	},
	{
		ID:         4,
		Name:       "iPhone 12 Pro",
		MerchantID: 1,
		Amount:     39000.00,
		Stock:      50,
	},
	{
		ID:         5,
		Name:       "iPhone 1 Pro",
		MerchantID: 1,
		Amount:     34000.00,
		Stock:      50,
	},
}

type fakeRepo struct{}

func (f fakeRepo) FindAllByMerchantID(MerchantID int) ([]models.Product, error) {
	var products []models.Product
	for _, p := range fakeProduct {
		if p.MerchantID == MerchantID {
			products = append(products, p)
		}
	}
	return products, nil
}

func (f fakeRepo) CreateProduct(product models.Product) (models.Product, error) {
	if len(fakeProduct) != 0 {
		lastProduct := fakeProduct[len(fakeProduct)-1]
		product.ID = lastProduct.ID + 1
	} else {
		product.ID = 1
	}
	fakeProduct = append(fakeProduct, product)
	return product, nil
}

func TestService_GetAllByMerchantID(t *testing.T) {
	tcs := []struct {
		testID      int
		description string
		merchantID  int
		wantLen     int
	}{
		{
			testID:      1,
			description: "Get products of merchant 1",
			merchantID:  1,
			wantLen:     4,
		},
		{
			testID:      2,
			description: "Get products of merchant 2",
			merchantID:  2,
			wantLen:     1,
		},
	}

	srv := NewProductService(fakeRepo{})
	for _, tc := range tcs {
		data, err := srv.GetAllByMerchantID(tc.merchantID)
		if err != nil {
			t.Error("found error")
			continue
		}
		if len(data) != tc.wantLen {
			t.Errorf("want length %d, but got %d", tc.wantLen, len(data))
		}
	}
}

func TestService_CreateProduct(t *testing.T) {
	tcs := []struct {
		testID      int
		description string
		merchantID  int
		product     models.Product
		wantProduct models.Product
		wantErr     string
	}{
		{
			testID:      1,
			description: "Create product for merchant 1",
			merchantID:  1,
			product: models.Product{
				Name:   "iPhone 13",
				Amount: 54050,
				Stock:  1,
			},
			wantProduct: models.Product{
				ID:         6,
				MerchantID: 1,
				Name:       "iPhone 13",
				Amount:     54050,
				Stock:      1,
			},
			wantErr: "",
		},
		{
			testID:      2,
			description: "Create product for merchant 1 but product already reached max items",
			merchantID:  2,
			product: models.Product{
				Name:   "iPhone 12",
				Amount: 100,
				Stock:  10,
			},
			wantProduct: models.Product{},
			wantErr:     "409",
		},
	}

	srv := NewProductService(fakeRepo{})
	for _, tc := range tcs {
		catched := false
		data, err := srv.CreateProduct(tc.merchantID, tc.product)
		if err != nil {
			if tc.testID == 2 && err.Error() == "409" {
				catched = true
				continue
			}
			t.Error("found error")
			continue
		}
		if tc.wantProduct != data && catched {
			t.Errorf("TEST ID %d: want %+v, but got %+v", tc.testID, tc.wantProduct, data)
		}
	}
}
