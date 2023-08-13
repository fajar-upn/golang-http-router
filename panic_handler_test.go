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
 * Panic Handler is useful for handling  panic error
 * automatically will occur error, and our web will be stop
 * sending response
 *
 * sometimes when panic occur, we want do something, for example
 * we want to tell mistake in website, or maybe send mistake information log
 * that occured
 *
 * Previously, like which already in golang website we want to panic handling
 * we must create spesific Middleware manually
 *
 * But in the Router already has for panic handler, The method with use
 * attribute PanicHandler: func(http.ResponseWriter, *http.Request, interface{})
 */

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(w, "panic :", error)
	}

	// note: without "router.PanicHandler" website will be occur page "this page isn't working"
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "panic :Ups", string(body))
}
