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
 * Route Pattern
 *
 * this is kind of Route Pattern:
 *
 * 1. Named Parameter
 * Namad Parameter is a pattern to create parameter with name,
 * every parameter name must prefix with ":" (colon), after that
 * followed with parameter name
 *
 * Example, if we have pattern like this:
 * |=============================================|
 * |     Pattern        |     /user/:user        |
 * |=============================================|
 * |    /user/eko       |        match           |
 * |    /user/you       |        match           |
 * |  /user/eko/profile |       no match         |
 * |      /user/        |       no match         |
 * |=============================================|
 */

func TestRoutePatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id/item/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		itemId := p.ByName("itemId")
		text := "Product " + id + " Item " + itemId

		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/product/1/item/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 1", string(body))
}

/**
 * 2. Catch All Parameter
 * besides named parameter, Route Pattern has catch all parameter, that is
 * catch all parameter
 *
 * Catch all parameter must prefix with * (asterisk), after that followed
 * with parameter name
 *
 * Catch all parameter must be at the end last of URL
 */
func TestRouteCatchAllNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/image/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		text := "Image: " + image

		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/image/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image: /small/profile.png", string(body))
}
