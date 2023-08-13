package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

/**
 * Serve File
 * In Golang Web Material, we already explain about Serve File
 * Also in Router support serve static file use function ServeFile
 * (Path, FileSystem)
 *
 * Where in the path, we must use Catch All Parameter
 * But in the Filesystem we can do manual load from folder or
 * use golang embed, like we have discussed in Golang Web
 */

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()

	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:8080/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request) //router filled with testing variable

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello HTTP Router", string(body))
}
