package services

import (
	"fmt"
	"go-api/internal/models"
	"go-api/internal/repositories"
	"go-api/internal/utils"

	"github.com/Pallinder/go-randomdata"
)

// MerchantService interface
type MerchantService interface {
	GetMerchantByID(ID int) (models.Merchant, error)
	CreateMerchant(name string, bankAccount string) (models.Merchant, error)
	UpdateMerchant(merchant models.Merchant) (models.Merchant, error)
}

// MerchantServ struct
type MerchantServ struct {
	Repo repositories.MerchantRepository
}

// NewMerchantService func
func NewMerchantService(repo repositories.MerchantRepository) *MerchantServ {
	return &MerchantServ{Repo: repo}
}

// CreateMerchant func
func (s MerchantServ) CreateMerchant(name string, bankAccount string) (models.Merchant, error) {
	result, _ := s.Repo.FindMerchantByBankAcc(bankAccount)
	if (models.Merchant{}) != result {
		return models.Merchant{}, utils.NewConflictError()
	}
	merchant := models.Merchant{
		Name:         name,
		BangkAccount: bankAccount,
		Username:     randomdata.Email(),
		Password:     utils.GeneratePassword(),
	}
	return s.Repo.CreateMerchant(merchant)
}

// UpdateMerchant func
func (s MerchantServ) UpdateMerchant(merchant models.Merchant) (models.Merchant, error) {
	updated, _ := s.Repo.UpdateMerchant(merchant)
	fmt.Printf("%+v \n", updated)
	if (models.Merchant{}) == updated {
		return models.Merchant{}, utils.NewNotFoundError()
	}
	return updated, nil
}

// GetMerchantByID func
func (s MerchantServ) GetMerchantByID(ID int) (models.Merchant, error) {
	return s.Repo.FindMerchantByID(ID)
}
