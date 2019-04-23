package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/pinpt/go-dremioproxy"
)

func main() {
	conn, err := sql.Open("dremioproxy", "http://localhost:9080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	rows, err := conn.QueryContext(context.Background(), `SELECT NOW()`)
	if err != nil {
		log.Fatal(err)
	}
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("columns:", cols)
	defer rows.Close()
	var count int
	started := time.Now()
	for rows.Next() {
		var val sql.NullFloat64
		if err := rows.Scan(&val); err != nil {
			log.Fatal(err)
		}
		if val.Valid {
			fmt.Println("val:", val)
		}
		count++
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	fmt.Println("count", count, "duration", time.Since(started))
}
