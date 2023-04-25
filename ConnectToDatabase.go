package main

import (
	"database/sql"
	"fmt"
	"strings"
	_ "time"

	_ "github.com/lib/pq"
)

func Connect(data CurrencyData) error {
	// открываем соединение с базой данных PostgreSQL
	connStr := "postgres://postgres:LetDoRehcfx@localhost/practice_database?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
	defer db.Close()
	rates := make(map[string]float64)
	for k, v := range data.Rates {
		rates[strings.ToLower(k)] = v
	}

	query := fmt.Sprintf(`
        INSERT INTO practice_schema.currency_data (base, date, success, timestamp, usd_rate, eur_rate, gbp_rate)
        VALUES ('%s', '%s', %t, %d, %f, %f, %f)
    `,
		data.Base,
		data.Date,

		data.Success,
		data.TimeStamp,
		rates["usd"],
		rates["eur"],
		rates["gbp"],
	)

	_, err = db.Exec(query)
	return err
}
