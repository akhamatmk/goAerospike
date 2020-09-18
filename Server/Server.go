package server

import (
	"github.com/akhamatvarokah/goAerospike/route"
	"github.com/labstack/echo/middleware"
)

// Run ...
func Run() {
	e := route.Route()
	e.Use(middleware.CORS())

	// Start server
	e.Start(":1323")
}
