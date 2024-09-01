package main

import (
	"net/http"

	middlweare "github.com/mnsdojo/custom-router-go/internal/middleware"
	"github.com/mnsdojo/custom-router-go/internal/router"
)

func main() {
	r := router.NewRouter()

	r.Use(middlweare.LoggerMiddleware)
	// Define your routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Home Page!"))
	})

	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Us"))
	})
	http.ListenAndServe(":8080", r)
}
