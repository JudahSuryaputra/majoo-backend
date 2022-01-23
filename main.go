package main

import (
	"log"
	"majoo-backend/cfg"
	"majoo-backend/http"
	"majoo-backend/repositories"

	_ "github.com/lib/pq"
)

func init() {
	cfg.Init()
}

func main() {
	app := http.App{}

	conn, err := repositories.Conn()
	if err != nil {
		log.Fatalf("Cannot initialize connection to database: %v", err)
	}

	app.Initialize(conn)
	app.RunServer()
}
