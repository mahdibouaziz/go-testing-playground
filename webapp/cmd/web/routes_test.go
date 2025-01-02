package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_application_routes(t *testing.T) {
	// verify all the routes all registred correctly
	var registred = []struct {
		route  string
		method string
	}{
		{route: "/", method: "GET"},
		{route: "/static/*", method: "GET"},
		{route: "/login", method: "POST"},
	}

	var app application
	mux := app.routes()
	chiRoutes := mux.(chi.Routes)

	for _, route := range registred {
		// check to see if the route exist
		if !routExists(route.route, route.method, chiRoutes) {
			t.Errorf("route %s is not registred", route.route)
		}

	}

}

func routExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false
	_ = chi.Walk(chiRoutes, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})

	return found
}
