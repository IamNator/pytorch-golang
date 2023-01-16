package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	
	"github.com/go-chi/chi"
	"gonum.org/v1/gonum/tensor"
	jose "gopkg.in/square/go-jose.v2"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/predict", func(w http.ResponseWriter, r *http.Request) {
		// Get input data from the request body
		var input interface{}
		// Apply the model to the input data
		var output interface{}
		// Write the output to the response body
		w.Write(output)
	})

	http.ListenAndServe(":3000", r)
}

