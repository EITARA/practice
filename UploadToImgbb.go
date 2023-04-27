package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func UpladImg() {
	// Открываем файл для отправки
	file, err := os.Open("G:\\exchange_rates.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Создаем "multipart/form-data" запрос
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer.Close()

	// Отправляем запрос на imgbb API
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.imgbb.com/1/upload?key=83fd5ca0969689a2076917c8fbbe7435", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

}
