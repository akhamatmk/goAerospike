package aerospikecontroller

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetAllBins ...
func GetAllBins(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
