package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"pet-system/interface/controllers"
)

func NewRouter(e *echo.Echo, c controllers.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pets", func(context echo.Context) error { return c.Pet.GetPet(context) })

	return e
}
