package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	p "github.com/pinpt/go-dremioproxy"
)

type debug struct {
}

func (d *debug) QueryFinished(query string, duration time.Duration, err error) {
	fmt.Printf("[DEBUG] query: %s, took %v, err: %v\n", query, duration, err)
}

func main() {
	conn, err := sql.Open("dremioproxy", "http://localhost:9080?skip-verify=true")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	p.SetDebug(&debug{})
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
		var val sql.NullString
		if err := rows.Scan(&val); err != nil {
			log.Fatal(err)
		}
		fmt.Println("val:", val)
		count++
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	fmt.Println("count", count, "duration", time.Since(started))
}
