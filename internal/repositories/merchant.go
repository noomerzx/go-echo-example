package repositories

import (
	"go-api/internal/app"
	"go-api/internal/models"
)

// MerchantRepository interface
type MerchantRepository interface {
	FindMerchantByID(ID int) (models.Merchant, error)
	FindMerchantByBankAcc(bankAcc string) (models.Merchant, error)
	CreateMerchant(merchant models.Merchant) (models.Merchant, error)
	UpdateMerchant(merchant models.Merchant) (models.Merchant, error)
}

// MerchantRepo struct
type MerchantRepo struct {
	cf *app.Config
}

// NewMerchantRepo func
func NewMerchantRepo(cf *app.Config) *MerchantRepo {
	return &MerchantRepo{cf: cf}
}

// CreateMerchant func
func (r MerchantRepo) CreateMerchant(merchant models.Merchant) (models.Merchant, error) {
	if len(r.cf.Db.Merchants) != 0 {
		lastMerchant := r.cf.Db.Merchants[len(r.cf.Db.Merchants)-1]
		merchant.ID = lastMerchant.ID + 1
	} else {
		merchant.ID = 1
	}
	r.cf.Db.Merchants = append(r.cf.Db.Merchants, merchant)
	return merchant, nil
}

// UpdateMerchant func
func (r MerchantRepo) UpdateMerchant(merchant models.Merchant) (models.Merchant, error) {
	var m models.Merchant
	for i := 0; i < len(r.cf.Db.Merchants); i++ {
		if r.cf.Db.Merchants[i].Username == merchant.Username && r.cf.Db.Merchants[i].Password == merchant.Password {
			r.cf.Db.Merchants[i].Name = merchant.Name
			m = r.cf.Db.Merchants[i]
		}
	}
	return m, nil
}

// FindMerchantByID func
func (r MerchantRepo) FindMerchantByID(ID int) (models.Merchant, error) {
	var merchant models.Merchant
	for _, m := range r.cf.Db.Merchants {
		if m.ID == ID {
			merchant = m
		}
	}
	return merchant, nil
}

// FindMerchantByBankAcc func
func (r MerchantRepo) FindMerchantByBankAcc(bankAcc string) (models.Merchant, error) {
	var merchant models.Merchant
	for _, m := range r.cf.Db.Merchants {
		if m.BangkAccount == bankAcc {
			merchant = m
		}
	}
	return merchant, nil
}
