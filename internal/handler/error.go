package handler

import (
	"fmt"
	"log"
	"net/http"
)

func buildErrorHandler(err interface{}, message string, statusCode int) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO: Add HTML error handler using error.html instead of plain text
		log.Fatal(err)
		res.Header().Set("Content-Type", "text/plain; charset=utf-8")
		res.Header().Set("X-Content-Type-Options", "nosniff")
		res.WriteHeader(statusCode)
		fmt.Fprintln(res, message)
	}
}
