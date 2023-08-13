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
 * Middleware
 * HTTPRouter it just library for http router, that is not has another features
 * other than Router
 *
 * Because Router is implementation from http.Handler, so we can create own middleware
 * as already discussed in Golang Web
 */

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive Request")
	middleware.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Middleware")
	})

	middleware := LogMiddleware{router}

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(body))
}
