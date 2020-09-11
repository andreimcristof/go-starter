package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// parsing an http response
	s := getContentAsString("https://dog.ceo/api/breeds/image/random")
	fmt.Println(s)

	res := asDogeStruct("https://dog.ceo/api/breeds/image/random")
	fmt.Println(res)
}

func getContentAsString(url string) string {
	// making the request here
	resp, err := http.Get(url)

	// we exit with a non-zero status if something failed with the request
	if err != nil {
		panic(err)
	}

	// deferred (scheduled) cleaning of resources once the current function ends
	defer resp.Body.Close()

	// we use the ReadAll method to retrieve the body of the response as array of bytes
	r, _ := ioutil.ReadAll(resp.Body)
	return string(r)
}

// DogeRes a dog
type DogeRes struct {
	Message string
	Status  string
}

func asDogeStruct(url string) DogeRes {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	// create a holder structure
	var d DogeRes

	// parse as struct by passing the array of bytes and a pointer to the target structure
	errU := json.Unmarshal(b, &d)

	if errU != nil {
		panic(nil)
	}
	return d
}
