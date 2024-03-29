package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_addIPToContext(t *testing.T) {

	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarder-For", "192.4.1.2", "", false},
		{"", "", "test:world", true},
	}

	var app application

	//create a dummy handler that we'll use to check the context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// make sure that the value exists in the context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}

		//make sure we got a string back
		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}

		t.Log(ip)

	})

	for _, e := range tests {
		// create the handler to test
		handlerToTest := app.addIPToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://testing", nil)

		if e.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)

	}

}
