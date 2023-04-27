package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ApiConnect() string {
	url := "https://api.apilayer.com/fixer/latest?symbols=GBP%2CEUR%2CUSD&base=RUB"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", "r6Ff7Q8GIowZr9rSUAscGlYC8xUtB5XC")

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)

	return string(body)
}
