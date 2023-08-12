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
 * Params
 *
 * httprouter.Handle has third parameter, which is Params.
 * Params is a variable for save parameter which send from client
 * But Params is not query parameter, but parameter in URL
 *
 * Sometimes we need craete dynamic URL, for an example "/product/1",
 * "/product/2", and others
 *
 * ServeMux doesn't support that things, but Router support dynamic URL
 *
 * Which dynamic parameter had in URL, automatically collected in Params
 * But, in order for Router to know, we must tell when we add Route,
 * Which part we can create a URL path to make dynamic
 */
func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		text := "Product " + id

		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/product/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))
}
