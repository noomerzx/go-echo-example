package handlers

import (
	"go-api/internal/app"
	"go-api/internal/middlewares"
	"go-api/internal/repositories"
	"go-api/internal/services"
	"net/http"

	"github.com/labstack/echo"
)

type route struct {
	HTTPMethod     string
	Endpoint       string
	HandlerFunc    echo.HandlerFunc
	MiddlewareFunc []echo.MiddlewareFunc
}

// NewRouter func
func NewRouter(e *echo.Echo, c *app.Config) error {
	merchantRepo := repositories.NewMerchantRepo(c)
	merchantService := services.NewMerchantService(merchantRepo)
	productRepo := repositories.NewProductRepo(c)
	productService := services.NewProductService(productRepo)
	reportService := services.NewReportService()
	merchantHandler := NewMerchantHandler(merchantService, productService, reportService)
	routes := []route{
		{
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/merchant/information/:id",
			HandlerFunc:    merchantHandler.GetMerchantByID,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.RequestHandlerMiddleware(c)},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/merchant/register",
			HandlerFunc:    merchantHandler.Register,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/merchant/update",
			HandlerFunc:    merchantHandler.Update,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.RequestHandlerMiddleware(c)},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/merchant/:id/product",
			HandlerFunc:    merchantHandler.CreateProduct,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.RequestHandlerMiddleware(c)},
		},
		{
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/merchant/:id/products",
			HandlerFunc:    merchantHandler.GetProducts,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.RequestHandlerMiddleware(c)},
		},
		{
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/merchant/:id/report",
			HandlerFunc:    merchantHandler.GenReport,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.RequestHandlerMiddleware(c)},
		},
	}

	for _, r := range routes {
		e.Add(r.HTTPMethod, "/api"+r.Endpoint, r.HandlerFunc, r.MiddlewareFunc...)
	}
	return nil
}
