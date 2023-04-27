package main

func main() {

	// выполнение запроса к API каждые 8 часов
	/*	interval := time.Hour * 8

		// Запуск функции, выполняющей запрос каждые 8 часов
		ticker := time.NewTicker(interval)
		for range ticker.C {*/
	jsonStr := ApiConnect()
	//}
	println("Complete api")
	parsingResult := Parsing(jsonStr)
	println("Complete parse")
	err := Connect(parsingResult)
	if err != nil {
		return
	}
	println("Complete conn")
	DrawingGraf()
	println("Complete dr")
	UpladImg()
	println("Complete upload")
	// Бесконечный цикл для продолжения работы приложения
	/*	for {
		time.Sleep(10 * time.Second)
	}*/
}
