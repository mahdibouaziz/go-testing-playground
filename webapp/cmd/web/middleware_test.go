package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_addIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{headerName: "", headerValue: "", addr: "", emptyAddr: false},
		{headerName: "", headerValue: "", addr: "", emptyAddr: true},
		{headerName: "X-Forwarded-For", headerValue: "192.3.2.1", addr: "", emptyAddr: false},
		{headerName: "", headerValue: "", addr: "Hello:invalid", emptyAddr: false},
	}

	var app application

	// create a dummy handler that we'll use to check the context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure that the ip exists in the context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Errorf("%s not present", contextUserKey)
		}

		// make sure we got a string back
		ip, ok := val.(string)
		if !ok {
			t.Errorf("%s not string", val)
		}
		t.Log(ip)
	})

	for _, val := range tests {
		// create the handler to test
		handlerToTest := app.addIPToContext(nextHandler)

		// create dummy request
		req := httptest.NewRequest("GET", "http://testing", nil)

		if val.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(val.headerName) > 0 {
			req.Header.Add(val.headerName, val.headerValue)
		}

		if len(val.addr) > 0 {
			req.RemoteAddr = val.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}

}

func Test_application_ipFromContext(t *testing.T) {
	//  create an app var of type application
	var app application

	// get a context
	ctx := context.Background()

	// put something in the context
	contextTestVal := "1.1.1.1"
	ctx = context.WithValue(ctx, contextUserKey, contextTestVal)

	// call the function
	result := app.ipFromContext(ctx)

	// perform the test
	if !strings.EqualFold(result, contextTestVal) {
		t.Errorf("wrong value returned from context, expectd %s, got %s", contextTestVal, result)
	}
}
