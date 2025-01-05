package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mahdibouaziz/go-testing-playground/webapp/pkg/repository"
	"github.com/mahdibouaziz/go-testing-playground/webapp/pkg/repository/dbrepo"
)

const port = 8090

type application struct {
	DSN       string
	DB        repository.DatabaseRepo
	Domain    string
	JWTSecret string
}

func main() {
	// should be env variables for prod, here this is just a playground
	app := application{
		DSN:       "postgres://postgres:postgres@localhost:5432/users",
		Domain:    "example.com",
		JWTSecret: "SecretTop",
	}

	// Connect to the db
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}
	defer conn.Close()
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	// print out a message
	log.Printf("Starting server on port %d \n", port)
	// start the server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
