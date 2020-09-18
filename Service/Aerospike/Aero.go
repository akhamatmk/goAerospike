package aerospike

import (
	"errors"
	"fmt"

	as "github.com/aerospike/aerospike-client-go"
)

type PaylodAerospike struct {
	NameSpace string      `json:"namespace" form:"name" query:"namespace"`
	SetName   string      `json:"setname" form:"setname" query:"setname"`
	Key       string      `json:"key" form:"key" query:"key"`
	Value     interface{} `json:"value" form:"value" query:"value"`
}

// MyStruct ...
type MyStruct struct {
	Key   string
	Value interface{}
}

// GetAllData ...
func GetAllData(nameSpace, setName string) []*MyStruct {
	result := []*MyStruct{}
	c, key, exists, err := checkExist(nameSpace, setName)
	if err != nil || key == nil || exists == false {
		return nil
	}

	stmt := as.NewStatement(nameSpace, setName)
	rs, err := c.Query(nil, stmt)
	if err == nil {

		var data map[string]interface{}
		for res := range rs.Results() {
			if res.Err != nil {
				fmt.Println("Err------", res.Err)
				return result
			}

			data = res.Record.Bins
		}

		for key, element := range data {
			result = append(result, &MyStruct{
				Key:   key,
				Value: element,
			})
		}
	}

	return result
}

func InsertData(data PaylodAerospike) *MyStruct {
	c, key, exists, err := checkExist(data.NameSpace, data.SetName)
	if err != nil || key == nil || exists == false {
		return nil
	}

	bins := as.BinMap{
		data.Key: data.Value,
	}

	err = c.Put(nil, key, bins)
	if err != nil {
		return nil
	}

	return &MyStruct{
		Key:   data.Key,
		Value: data.Value,
	}
}

// GetValueByKey ...
func GetValueByKey(nameSpace, setName, k string) interface{} {
	// define a client to connect to
	c, key, exists, err := checkExist(nameSpace, setName)
	if err != nil || key == nil || exists == false {
		return nil
	}

	rec, err := c.Get(nil, key)
	if err != nil {
		panic(err)
	}

	return rec.Bins[k]
}

func checkExist(nameSpace, setName string) (*as.Client, *as.Key, bool, error) {
	c := GetAerospikeClient()
	key, err := as.NewKey(nameSpace, setName, "key") // user key can be of any supported type
	if err != nil {
		return nil, nil, false, errors.New("Namespace not found")
	}

	exists, err := c.Exists(nil, key)
	if err != nil {
		return nil, nil, false, errors.New("Namespace not exist")
	}

	return c, key, exists, nil
}

func Edit(data PaylodAerospike) *MyStruct {
	c, key, exists, err := checkExist(data.NameSpace, data.SetName)
	if err != nil || key == nil || exists == false {
		return nil
	}

	rec, err := c.Get(nil, key)
	if err != nil {
		panic(err)
	}

	if rec.Bins[data.Key] == nil {
		return nil
	}

	bins := as.BinMap{
		data.Key: data.Value,
	}

	err = c.Put(nil, key, bins)
	if err != nil {
		return nil
	}

	return &MyStruct{
		Key:   data.Key,
		Value: data.Value,
	}
}

func Destroy(data PaylodAerospike) interface{} {
	c, key, exists, err := checkExist(data.NameSpace, data.SetName)
	if err != nil || key == nil || exists == false {
		return nil
	}

	err = c.Put(nil, key, as.BinMap{data.Key: nil})
	if err != nil {
		return "Failed delete data or data not exit"
	}

	return "Success Delete"
}

// GetAerospkeClient ..
func GetAerospikeClient() *as.Client {
	client, err := as.NewClient("172.28.128.4", 3000)
	if err != nil {
		panic(err)
	}

	return client
}
