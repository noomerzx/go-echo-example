package middlewares

import (
	"fmt"
	"go-api/internal/app"

	"go-api/internal/models"
	"go-api/internal/utils"

	"github.com/labstack/echo"
)

// RequestHandlerMiddleware func (Each Handler)
func RequestHandlerMiddleware(config *app.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			pass := false
			m := new(models.Merchant)
			if err := c.Bind(m); err != nil {
				return utils.JSONResponse(c, nil, nil)
			}
			for _, mdb := range config.Db.Merchants {
				if m.Username == mdb.Username && m.Password == mdb.Password {
					pass = true
				}
			}
			if pass {
				return next(c)
			} else {
				return utils.JSONResponse(c, nil, utils.NewUnauthorizedError())
			}
		}
	}
}

// RequestMiddleware func (Global)
func RequestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		fmt.Println("Requested in global middleware")
		return next(c)
	}
}
