package controllers

import (
	"net/http"

	ar "github.com/akhamatvarokah/goAerospike/service/aerospike"
	Utils "github.com/akhamatvarokah/goAerospike/utils"
	"github.com/labstack/echo"
)

type postData struct {
	NameSpace string      `json:"namespace" form:"namespace" query:"namespace"`
	SetName   string      `json:"setname" form:"setname" query:"setname"`
	Key       string      `json:"key" form:"key" query:"key"`
	Value     interface{} `json:"value" form:"value" query:"value"`
	KeyBin    string      `json:"keybin" form:"keybin" query:"keybin"`
}

func GetNameSpace(c echo.Context) error {
	setName := [2]string{"knox", "mnemonic"}
	return c.JSON(http.StatusOK, Utils.ResponseOk(setName))
}

func GetSetName(c echo.Context) error {
	namespace := c.Param("namespace")

	data := ar.GetAllSetname(namespace)
	return c.JSON(http.StatusOK, Utils.ResponseOk(data))
}

// Getdata ...
func Getdata(c echo.Context) error {
	key := c.Param("key")
	namespace := c.QueryParam("namespace")
	setname := c.QueryParam("setname")
	keybin := c.QueryParam("key_bin")
	filter := c.QueryParam("filter")

	if key == "" {
		data := ar.GetAllData(namespace, setname, keybin, filter)

		if len(data) > 0 {
			return c.JSON(http.StatusOK, Utils.ResponseOk(data))
		} else {
			return c.JSON(http.StatusOK, Utils.ResponseOk(nil))
		}

	} else {
		result := ar.GetValueByKey(namespace, setname, keybin, key)
		return c.JSON(http.StatusOK, Utils.ResponseOk(result))
	}
}

func Edit(c echo.Context) error {
	u := new(postData)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if u.Key == "" || u.Value == "" {
		return c.JSON(http.StatusBadRequest, "Key or value cant null")
	}

	if u.NameSpace == "" {
		return c.JSON(http.StatusBadRequest, "NameSpace or value cant null")
	}

	if u.SetName == "" {
		return c.JSON(http.StatusBadRequest, "SetName or value cant null")
	}

	if u.KeyBin == "" {
		return c.JSON(http.StatusBadRequest, "Bin or value cant null")
	}

	result := ar.Edit(ar.PaylodAerospike{
		NameSpace: u.NameSpace,
		SetName:   u.SetName,
		Key:       u.Key,
		Value:     u.Value,
	})

	return c.JSON(http.StatusOK, result)
}

func DeleteData(c echo.Context) error {
	key := c.Param("key")
	namespace := c.Param("namespace")
	setname := c.Param("setname")
	bin := c.Param("bin")

	result := ar.Destroy(ar.PaylodAerospike{
		NameSpace: namespace,
		SetName:   setname,
		Key:       key,
		KeyBin:    bin,
	})

	return c.JSON(http.StatusOK, Utils.ResponseOk(result))
}

// Post ...
func Insert(c echo.Context) error {
	u := new(postData)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if u.Key == "" || u.Value == "" {
		return c.JSON(http.StatusBadRequest, "Key or value cant null")
	}

	result := ar.InsertData(ar.PaylodAerospike{
		NameSpace: u.NameSpace,
		SetName:   u.SetName,
		Key:       u.Key,
		Value:     u.Value,
		KeyBin:    u.KeyBin,
	})
	return c.JSON(http.StatusOK, Utils.ResponseOk(result))
}

