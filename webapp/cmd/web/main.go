package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type application struct {
	Session *scs.SessionManager
	DSN     string
	DB      *pgxpool.Pool
}

func main() {
	// set up an app config
	app := application{
		Session: getSession(),
	}

	// Connect to the db
	app.DSN = "postgres://postgres:postgres@localhost:5432/users"
	// app.DSN = os.Getenv("DATABASE_URL") // TODO - change this later
	db, err := app.connectToDB()
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}
	defer db.Close() // Ensure the database connection is closed on exit
	app.DB = db

	// print out a message
	log.Println("Starting server on port 8080")

	// start the server
	err = http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
