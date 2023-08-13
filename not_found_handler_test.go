package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

/**
 * Not Found Handler
 *
 * Besides Panic Handler, Router also has not found handler.
 * Not Found Handler is a handler which will executed when client
 * try do a request URL which doesn't exit in the Router
 *
 * Defaultly, if route not found, Router will be direct request
 * to http.NotFound. but we can change it
 *
 * The method is change router.NotFound to http.Handler
 */

func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Not Found")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Not Found", string(body))
}
