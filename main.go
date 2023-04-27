package main

func main() {

	// выполнение запроса к API каждые 8 часов
	/*	interval := time.Hour * 8

		// Запуск функции, выполняющей запрос каждые 8 часов
		ticker := time.NewTicker(interval)
		for range ticker.C {*/
	jsonStr := ApiConnect()
	//}
	parsingResult := Parsing(jsonStr)
	err := Connect(parsingResult)
	if err != nil {
		return
	}
	DrawingGraf()
	// Бесконечный цикл для продолжения работы приложения
	/*	for {
		time.Sleep(10 * time.Second)
	}*/
}
