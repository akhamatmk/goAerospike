package http

import (
	"net/http"

	aec "github.com/akhamatvarokah/goAerospike/api/contollers/Aerospikecontroller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Run ...
func Run() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", aec.GetAllBins)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
