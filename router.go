package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/**
 * Router
 * Core from library HTTPRputer is struct Router
 *
 * This router is implementation from http.Handler, so we can easily
 * add to the http.Server
 *
 * for create Router, we can use function httprouter.New(), which will
 * return Router pointer
 */

// func main() {
// 	router := httprouter.New()

// 	server := http.Server{
// 		Handler: router,
// 		Addr:    "localhost:8080",
// 	}

// 	server.ListenAndServe()
// }

/**
 * HTTP Method
 *
 * Router similar with the serverMux, where we can add route in the Router
 *
 * The advantage when we compare with serveMux is we can organize HTTP method
 * in the Router (GET, POST, PUT and many more)
 *
 * how to add route in the Router is we can use which same function with their
 * HTTP method, for the example router.GET(), router.POST(), and others
 */

/**
 * httprouter.Handle
 *
 * When we use ServeMux, we can add route, we can add http.Handler
 *
 * different with Router, in the Router we can user http.Handler again,
 * but use type httprouter.Handle
 *
 * the main different with http.Handler is, in httprouter.Handlem has third parameter
 * that is Params, which will explain in another chapter
 * syntax:
 * type Handle func(http.ResponseWriter, *http.Request, Params)
 */

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello Httprouter")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8000",
	}

	server.ListenAndServe()
}
