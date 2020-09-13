package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	ac "github.com/akhamatvarokah/goAerospike/controllers/aerospikecontroller"
)

// Route ...
func Route() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", ac.Getdata)
	e.GET("/:key", ac.Getdata)

	e.POST("/", ac.Post)

	return e
}
