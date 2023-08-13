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
 * Method Not Allowed Handler
 *
 * When we use ServeMux, we cannot choose HTTP Method will be used
 * for Handler
 *
 * but in Router, we can determine HTTP Method which will be used,
 * So what happen when client not send HTTP method appropriate with
 * determine method ?
 *
 * So will be error occured "Method not Allowed"
 *
 * Defaultly, if error occur like that, th Router will be called function
 * http.Error
 *
 * If we want to change that, we can use router.MethodNotAllowed = http.Handler
 */

func TestNotAllowedHandler(t *testing.T) {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Method Not Allowed")
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "POST")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Method Not Allowed", string(body))
}
