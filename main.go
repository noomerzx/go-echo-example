package main

import (
	"flag"
	"fmt"
	"go-api/internal/app"
	"go-api/internal/handlers"
	"go-api/internal/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	env := flag.String("env", "dev", "environment")
	flag.Parse()
	c := app.NewConfig(*env)
	if err := c.Init(); err != nil {
		fmt.Println(err)
	}
	e := echo.New()
	if err := handlers.NewRouter(e, c); err != nil {
		fmt.Println("New Router Failed.")
	}
	e.Use(middleware.Logger())
	e.Use(middlewares.RequestMiddleware)
	e.Logger.Fatal(e.Start(":1323"))
}
