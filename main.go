package main

import (
	"encoding/base32"
	"encoding/json"
	"fmt"
	"time"
)

const (
	a = "a"
	b = "b"
	c = "c"
)


type App struct {
	Id string `json:"id"`
}

type Org struct {
	Name string `json:"name"`
}

type AppWithOrg struct {
	App
	Org
}

func main() {
	type FruitBasket struct {
		Name    string
		Fruit   []string
		Id      int64 `json:"ref"`// 声明对应的json key
		Created time.Time
	}

	jsonData := []byte(`
    {
        "Name": "Standard",
        "Fruit": [
             "Apple",
            "Banana",
            "Orange"
        ],
        "ref": 999,
        "Created": "2018-04-09T23:00:00Z"
    }`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", basket)
	fmt.Println(basket.Name, basket.Fruit, basket.Id)
	fmt.Println(basket.Created)
}


type S interface {
	printA()
	printB()
}

type tt func() S

type Struct struct {
	A struct {
	    SDASD func(string) (string) `perm:"read"`
    }

}

var Test map[string][]string

const encodeStd = "abcdefghijklmnopqrstuvwxyz234567"

// AddressEncoding defines the base32 config used for address encoding and decoding.
var AddressEncoding = base32.NewEncoding(encodeStd)
