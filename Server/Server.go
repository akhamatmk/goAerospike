package server

import (
	"github.com/akhamatvarokah/goAerospike/route"
)

// Run ...
func Run() {
	e := route.Route()
	// Start server
	e.Start(":1323")
}
