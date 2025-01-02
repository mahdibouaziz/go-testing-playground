package main

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/mahdibouaziz/go-testing-playground/webapp/pkg/data"
	"github.com/mahdibouaziz/go-testing-playground/webapp/pkg/db"
)

type application struct {
	Session *scs.SessionManager
	DSN     string
	DB      db.PostgresConn
}

func main() {
	gob.Register(data.User{})
	// set up an app config
	app := application{
		Session: getSession(),
	}

	// Connect to the db
	app.DSN = "postgres://postgres:postgres@localhost:5432/users"
	// app.DSN = os.Getenv("DATABASE_URL") // TODO - change this later if you need a dynamic db (for this example it is just a playground)
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}
	defer conn.Close()
	app.DB = db.PostgresConn{DB: conn}

	// print out a message
	log.Println("Starting server on port 8080")

	// start the server
	err = http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
