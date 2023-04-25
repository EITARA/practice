package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type CurrencyData struct {
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
	Success   bool               `json:"success"`
	TimeStamp int64              `json:"timestamp"`
}

func Parsing() CurrencyData {
	file, err := os.Open("G:\\go\\dippractice\\json\\test.json")
	if err != nil {
		fmt.Println("Could not open file:", err)
		return CurrencyData{}
	}
	defer file.Close()

	var currencyData CurrencyData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&currencyData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return CurrencyData{}
	}

	fmt.Println("Base currency:", currencyData.Base)
	fmt.Println("Date:", currencyData.Date)
	fmt.Println("Success:", currencyData.Success)
	fmt.Println("Time stamp:", currencyData.TimeStamp)
	fmt.Println("Rates:")
	for currency, rate := range currencyData.Rates {
		fmt.Printf("\t%s: %.2f\n", currency, rate)
	}
	return currencyData
}
