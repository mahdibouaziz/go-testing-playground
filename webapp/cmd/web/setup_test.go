package main

import (
	"log"
	"os"
	"testing"

	"github.com/mahdibouaziz/go-testing-playground/webapp/pkg/db"
)

var app application

func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	app.Session = getSession()

	// Connect to the db
	app.DSN = "postgres://postgres:postgres@localhost:5432/users"
	// app.DSN = os.Getenv("DATABASE_URL") // TODO - change this later if you need a dynamic db (for this example it is just a playground)
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}
	defer conn.Close()
	app.DB = db.PostgresConn{DB: conn}

	os.Exit(m.Run())
}
