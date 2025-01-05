package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (app *application) connectToDB() (*pgxpool.Pool, error) {
	// Get the database connection string from an environment variable
	// dsnExample := "postgres://username:password@localhost:5432/database_name"
	dsn := app.DSN
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Create a context with a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to the database
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = dbpool.Ping(ctx)
	if err != nil {
		dbpool.Close()
		return nil, err
	}

	log.Println("Connected to the database")
	return dbpool, nil
}
