package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
}

func main() {
	// set application config
	var app application

	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5433 user=postgres password=password dbname=movies sslmode=disable timezone=Asia/Tokyo connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to database

	app.Domain = "example.com"

	// start a web server
	fmt.Printf("Starting application on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
