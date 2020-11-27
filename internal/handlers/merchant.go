package handlers

import (
	"go-api/internal/models"
	"go-api/internal/services"
	"go-api/internal/utils"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/labstack/echo"
)

// Handler struct
type Handler struct {
	mService services.MerchantService
	pService services.ProductService
	rService services.ReportService
}

// NewMerchantHandler func
func NewMerchantHandler(mService services.MerchantService, pService services.ProductService, rService services.ReportService) *Handler {
	return &Handler{mService, pService, rService}
}

// GetMerchantByID func
func (h Handler) GetMerchantByID(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.JSONResponse(c, nil, nil)
	}
	result, err := h.mService.GetMerchantByID(ID)
	if err != nil {
		return utils.JSONResponse(c, nil, err)
	}
	return utils.JSONResponse(c, result, nil)
}

// Register func
func (h Handler) Register(c echo.Context) error {
	m := new(models.Merchant)
	if err := c.Bind(m); err != nil {
		return utils.JSONResponse(c, nil, nil)
	}
	result, err := h.mService.CreateMerchant(m.Name, m.BangkAccount)
	if err != nil {
		return utils.JSONResponse(c, nil, err)
	}
	return utils.JSONResponse(c, result, nil)
}

// Update func
func (h Handler) Update(c echo.Context) error {
	m := new(models.Merchant)
	if err := c.Bind(m); err != nil {
		return utils.JSONResponse(c, nil, nil)
	}
	result, err := h.mService.UpdateMerchant(*m)
	if err != nil {
		return utils.JSONResponse(c, nil, err)
	}
	return utils.JSONResponse(c, result, nil)
}

// CreateProduct func
func (h Handler) CreateProduct(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.JSONResponse(c, nil, nil)
	}
	p := new(models.Product)
	if err := c.Bind(p); err != nil {
		return utils.JSONResponse(c, nil, nil)
	}
	product, err := h.pService.CreateProduct(ID, *p)
	if err != nil {
		return utils.JSONResponse(c, nil, err)
	}
	result := models.ProductResponse{
		ID:     product.ID,
		Name:   product.Name,
		Amount: humanize.Commaf(product.Amount),
		Stock:  humanize.Comma(int64(product.Stock)),
	}
	return utils.JSONResponse(c, result, nil)
}

// GetProducts func
func (h Handler) GetProducts(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.JSONResponse(c, nil, nil)
	}
	result, err := h.pService.GetAllByMerchantID(ID)
	if err != nil {
		return utils.JSONResponse(c, nil, err)
	}
	return utils.JSONResponse(c, result, nil)
}

// GenReport func
func (h Handler) GenReport(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.JSONResponse(c, nil, nil)
	}
	h.rService.GenReport(ID, "")
	return utils.JSONResponse(c, nil, nil)
}
