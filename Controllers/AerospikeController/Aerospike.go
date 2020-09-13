package aerospikecontroller

import (
	"net/http"

	ar "github.com/akhamatvarokah/goAerospike/service/aerospike"

	"github.com/labstack/echo"
)

type postData struct {
	NameSpace string      `json:"namespace" form:"name" query:"namespace"`
	SetName   string      `json:"setname" form:"setname" query:"setname"`
	Key       string      `json:"key" form:"key" query:"key"`
	Value     interface{} `json:"value" form:"value" query:"value"`
}

// Getdata ...
func Getdata(c echo.Context) error {
	key := c.Param("key")
	if key == "" {

		data := ar.GetAllData("test", "aerospike")

		if len(data) > 0 {
			return c.JSON(http.StatusOK, data)
		} else {
			return c.JSON(http.StatusOK, nil)
		}

	} else {
		result := ar.GetValueByKey("test", "aerospike", key)
		return c.JSON(http.StatusOK, result)
	}
}

// Post ...
func Post(c echo.Context) error {
	u := new(postData)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if u.Key == "" || u.Value == "" {
		return c.JSON(http.StatusBadRequest, "Key or value cant null")
	}

	if u.NameSpace == "" {
		u.NameSpace = "test"
	}

	if u.SetName == "" {
		u.SetName = "aerospike"
	}

	result := ar.InsertData(ar.PaylodAerospike{
		NameSpace: u.NameSpace,
		SetName:   u.SetName,
		Key:       u.Key,
		Value:     u.Value,
	})
	return c.JSON(http.StatusOK, result)
}
