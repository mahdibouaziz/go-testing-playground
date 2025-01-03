package main

import (
	"os"
	"testing"

	"github.com/mahdibouaziz/go-testing-playground/webapp/pkg/repository/dbrepo"
)

var app application

func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	app.Session = getSession()
	// repo for testing without the db
	app.DB = &dbrepo.TestDBRepo{}

	os.Exit(m.Run())
}
