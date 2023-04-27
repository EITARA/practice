package main

import (
	"database/sql"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"sort"
	_ "strings"
	"time"
)

func DrawingGraf() {
	connStr := "postgres://postgres:LetDoRehcfx@localhost/practice_database?sslmode=disable" //сделать авторизацию
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
	defer db.Close()
	// Получаем данные из базы данных
	rows, err := db.Query("SELECT timestamp AS date, usd_rate, eur_rate, gbp_rate FROM practice_schema.currency_data")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Создаем точки на основе полученных данных
	usdPoints := make(plotter.XYs, 0)
	eurPoints := make(plotter.XYs, 0)
	gbpPoints := make(plotter.XYs, 0)

	for rows.Next() {
		var date int64
		var usdRate, eurRate, gbpRate float64
		err := rows.Scan(&date, &usdRate, &eurRate, &gbpRate)
		if err != nil {
			log.Fatal(err)
		}

		t := time.Unix(date, 0)

		usdPoints = append(usdPoints, plotter.XY{X: float64(t.Unix()), Y: usdRate})
		eurPoints = append(eurPoints, plotter.XY{X: float64(t.Unix()), Y: eurRate})
		gbpPoints = append(gbpPoints, plotter.XY{X: float64(t.Unix()), Y: gbpRate})
	}

	// Создаем график
	p := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	//Сортировка по времени
	sortedUsdPoints := sortPoints(usdPoints)
	sortedEurPoints := sortPoints(eurPoints)
	sortedGbpPoints := sortPoints(gbpPoints)

	p.Title.Text = "Курсы Валют"
	p.X.Label.Text = "Дата и Время"
	p.Y.Label.Text = "Курс Валют"

	// Создаем графические объекты и добавляем их на график
	usdLine, err := plotter.NewLine(sortedUsdPoints)
	if err != nil {
		log.Fatal(err)
	}
	usdLine.Color = color.RGBA{R: 255, A: 255}
	p.Add(usdLine)

	eurLine, err := plotter.NewLine(sortedEurPoints)
	if err != nil {
		log.Fatal(err)
	}
	eurLine.Color = color.RGBA{B: 255, A: 255}
	p.Add(eurLine)

	gbpLine, err := plotter.NewLine(sortedGbpPoints)
	if err != nil {
		log.Fatal(err)
	}
	gbpLine.Color = color.RGBA{G: 255, A: 255}
	p.Add(gbpLine)
	// Устанавливаем формат времени для оси X
	ticks := plot.TimeTicks{Format: "2006-01-02\n15:04:05"}

	// Форматирование времени для подписей значений
	p.X.Tick.Label.Font.Size = vg.Points(9)

	// Устанавливаем формат времени для оси X
	p.X.Tick.Marker = ticks
	p.Legend.Add("USD", usdLine)
	p.Legend.Add("EUR", eurLine)
	p.Legend.Add("GBP", gbpLine)

	// Сохраняем график в файл
	if err := p.Save(10*vg.Inch, 5*vg.Inch, "G:\\exchange_rates.png"); err != nil {
		log.Fatal(err)
	}

}
func sortPoints(points plotter.XYs) plotter.XYs {
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})
	return points
}
