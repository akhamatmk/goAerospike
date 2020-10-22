package aerospike

import (
	"errors"
	"fmt"
	as "github.com/aerospike/aerospike-client-go"
	"strings"
)

type PaylodAerospike struct {
	NameSpace string      `json:"namespace" form:"name" query:"namespace"`
	SetName   string      `json:"setname" form:"setname" query:"setname"`
	Key       string      `json:"key" form:"key" query:"key"`
	Value     interface{} `json:"value" form:"value" query:"value"`
	KeyBin    string      `json:"key_bin" form:"key_bin" query:"key_bin"`
}

// MyStruct ...
type MyStruct struct {
	Key   string
	Value interface{}
	SetName string
}

// GetAllData ...
func GetAllData(nameSpace, setName, keyBin , filter string) []*MyStruct {
	var result []*MyStruct
	client := GetAerospikeClient()
	stmt := as.NewStatement(nameSpace, setName)
	rs, _ := client.Query(nil, stmt)

	for res := range rs.Results() {
		if res.Err == nil {
			fmt.Println(res.Record.Key.SetName())
			fmt.Println(res.Record.Key.Namespace())

			for k , v := range res.Record.Bins {

				if filter != ""  {
					if strings.Contains(k , filter) {
						result = append(result, &MyStruct{
							Key:   k,
							Value: v,
							SetName: res.Record.Key.SetName(),
						})
					}
				}else {
					result = append(result, &MyStruct{
						Key:   k,
						Value: v,
						SetName: res.Record.Key.SetName(),
					})
				}


			}
		}
	}

	return result
}

// contains ...
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// InsertData ...
func InsertData(data PaylodAerospike) *MyStruct {
	client := GetAerospikeClient()
	key, err := as.NewKey(data.NameSpace, data.SetName, data.KeyBin)
	if err != nil {
		return nil
	}

	bins := as.NewBin(data.Key, data.Value)
	_ = client.PutBins(nil, key, bins)

	return &MyStruct{
		Key:   data.Key,
		Value: data.Value,
		SetName: data.SetName,
	}
}

// GetAllSetname ...
func  GetAllSetname(namespace string)  *[]string{
	var data []string
	client := GetAerospikeClient()
	stmt := as.NewStatement(namespace, "")
	rs, _ := client.Query(nil, stmt)

	for res := range rs.Results() {
		if res.Err == nil {
			if ! contains(data , res.Record.Key.SetName()) {
				data = append(data, res.Record.Key.SetName())
			}
		}
	}

	return  &data
}

// GetValueByKey ...
func GetValueByKey(nameSpace, setName, keyBin, k string) interface{} {
	// define a client to connect to
	c, key, exists, err := checkExist(nameSpace, setName, keyBin)
	if err != nil || key == nil || exists == false {
		return nil
	}

	rec, err := c.Get(nil, key)
	if err != nil {
		panic(err)
	}

	return rec.Bins[k]
}

func checkExist(nameSpace, setName string, keyBin interface{}) (*as.Client, *as.Key, bool, error) {
	c := GetAerospikeClient()
	key, err := as.NewKey(nameSpace, setName, keyBin) // user key can be of any supported type
	if err != nil {
		return nil, nil, false, errors.New("namespace not found")
	}

	exists, err := c.Exists(nil, key)
	if err != nil {
		return nil, nil, false, errors.New("namespace not exist")
	}

	return c, key, exists, nil
}

func Edit(data PaylodAerospike) *MyStruct {
	c, key, exists, err := checkExist(data.NameSpace, data.SetName, data.KeyBin)
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
	c, key, exists, err := checkExist(data.NameSpace, data.SetName, data.KeyBin)
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
	client, err := as.NewClient("172.28.128.3", 3000)
	if err != nil {
		panic(err)
	}

	return client
}
