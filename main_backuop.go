package main

import (
	"fmt"

	aero "github.com/aerospike/aerospike-client-go"
)

func test() {
	// define a client to connect to
	client, err := aero.NewClient("172.28.128.4", 3000)
	panicOnError(err)

	namespace := "test"
	setName := "aerospike"
	key, err := aero.NewKey(namespace, setName, "key") // user key can be of any supported type
	panicOnError(err)

	// define some bins
	// bins := aero.BinMap{
	// 	"bin1": 42, // you can pass any supported type as bin value
	// 	"bin2": "An elephant is a mouse with an operating system",
	// 	"bin3": []interface{}{"Go", 17981},
	// }

	// write the bins
	// writePolicy := aero.NewWritePolicy(0, 0)
	// err = client.Put(writePolicy, key, bins)
	// panicOnError(err)

	readPolicy := aero.NewPolicy()
	exists, err := client.Exists(readPolicy, key)
	panicOnError(err)
	fmt.Printf("key exists: %#v\n", exists)

	rec2, err := client.Get(readPolicy, key)
	panicOnError(err)

	fmt.Printf("value of %s: %v\n", "bin1", rec2.Key)

	// read it back!
	// readPolicy := aero.NewPolicy()
	// rec, err := client.Get(readPolicy, key)
	// panicOnError(err)

	// fmt.Printf("%#v\n", *rec)

	// // Add to bin1
	// err = client.Add(writePolicy, key, aero.BinMap{"bin1": 1})
	// panicOnError(err)

	// rec2, err := client.Get(readPolicy, key)
	// panicOnError(err)

	// fmt.Printf("value of %s: %v\n", "bin1", rec2.Bins["bin1"])

	// // prepend and append to bin2
	// err = client.Prepend(writePolicy, key, aero.BinMap{"bin2": "Frankly:  "})
	// panicOnError(err)
	// err = client.Append(writePolicy, key, aero.BinMap{"bin2": "."})
	// panicOnError(err)

	// rec3, err := client.Get(readPolicy, key)
	// panicOnError(err)

	// fmt.Printf("value of %s: %v\n", "bin2", rec3.Bins["bin2"])

	// // delete bin3
	// err = client.Put(writePolicy, key, aero.BinMap{"bin3": nil})
	// rec4, err := client.Get(readPolicy, key)
	// panicOnError(err)

	// fmt.Printf("bin3 does not exist anymore: %#v\n", *rec4)

	// // check if key exists
	// exists, err := client.Exists(readPolicy, key)
	// panicOnError(err)
	// fmt.Printf("key exists in the database: %#v\n", exists)

	// // delete the key, and check if key exists
	// existed, err := client.Delete(writePolicy, key)
	// panicOnError(err)
	// fmt.Printf("did key exist before delete: %#v\n", existed)

	// exists, err = client.Exists(readPolicy, key)
	// panicOnError(err)
	// fmt.Printf("key exists: %#v\n", exists)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
