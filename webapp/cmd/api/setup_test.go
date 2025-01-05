package main

import (
	"os"
	"testing"

	"github.com/mahdibouaziz/go-testing-playground/webapp/pkg/repository/dbrepo"
)

var app application

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "SecretTop"
	os.Exit(m.Run())
}
