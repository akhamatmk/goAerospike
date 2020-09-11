package aerospike

import (
	"fmt"
	"strconv"
	"time"

	as "github.com/aerospike/aerospike-client-go"
)

var client *as.Client

// MyStruct ...
type MyStruct struct {
	Key   string
	Value interface{}
}

// GetAllData ...
func GetAllData(nameSpace, setName string) []*MyStruct {
	result := []*MyStruct{}

	stmt := as.NewStatement(nameSpace, setName)
	rs, err := GetAerospikeClient().Query(nil, stmt)
	if err == nil {

		var data map[string]interface{}
		for res := range rs.Results() {
			if res.Err != nil {
				// handle error here
				// if you want to exit, cancel the recordset to release the resources
				fmt.Println("Err------", res.Err)
				return result

			} else {
				data = res.Record.Bins
			}
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

// GetAerospikeClient ...
func GetAerospikeClient() *as.Client {
	var err error
	port, _ := strconv.Atoi("3000")
	maxconn, _ := strconv.Atoi("10")
	host := "172.28.128.4"
	timeout, _ := strconv.Atoi("50")
	idletimeout, _ := strconv.Atoi("3600")
	clientPolicy := as.NewClientPolicy()
	clientPolicy.ConnectionQueueSize = maxconn
	clientPolicy.LimitConnectionsToQueueSize = true
	clientPolicy.Timeout = time.Duration(timeout) * time.Millisecond
	clientPolicy.IdleTimeout = time.Duration(idletimeout) * time.Second
	client, err = as.NewClientWithPolicy(clientPolicy, host, port)
	if err != nil {
		panic(err)
	}

	return client
}
