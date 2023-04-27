package main

import (
	"fmt"
	"time"
)

func main() {

	// создадим новый ticker
	ticker := time.NewTicker(8 * time.Hour)
	defer ticker.Stop()

	// бесконечный цикл для выполнения кода
	for {
		select {
		// проверяем, что ticker двинулся
		case <-ticker.C:
			//код, который нужно выполнять каждые 8 часов
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
			fmt.Println("Приложение запущено через каждые 8 часов")
		}
	}

}
