package aerospikecontroller

import (
	"net/http"

	ar "github.com/akhamatvarokah/goAerospike/service/aerospike"
	"github.com/labstack/echo"
)

// Getdata ...
func Getdata(c echo.Context) error {

	data := ar.GetAllData("test", "aerospike")
	if len(data) > 0 {
		return c.JSON(http.StatusCreated, data)
	} else {
		return c.JSON(http.StatusCreated, nil)
	}
}

// GetdataByKey ...
func GetdataByKey(c echo.Context) error {
	key := c.Param("key")
	data := ar.GetValueByKey("test", "aerospike", key)

	return c.JSON(http.StatusCreated, data)
}
