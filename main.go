package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	parsingResult := Parsing()
	fmt.Println(parsingResult)
	err := Connect(parsingResult)
	if err != nil {
		return
	}
	DrawingGraf()
}
