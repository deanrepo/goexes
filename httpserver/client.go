package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {

	resp, err := http.Get("http://localhost:8090")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	fmt.Println("Response status code:", resp.StatusCode)

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("====>response body: %+v", strconv.Quote(string(responseBody)))
}
