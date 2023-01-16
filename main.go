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


//loading and running a PyTorch model
func pytorchModel() {
    // Load the PyTorch model
    var model *tensor.Dense

    // Run the model on input data
    input := tensor.NewDense(...)
    output := model.Apply(input)

    fmt.Println(output)
}


//encrypt a PyTorch model
func encrypt() {
    // Load the PyTorch model
    var model *tensor.Dense

    // Serialize the model
    modelJSON, err := json.Marshal(model)
    if err != nil {
        panic(err)
    }

    // Encrypt the model
    key := []byte("example key")
    encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.DIRECT, Key: key}, nil)
    if err != nil {
        panic(err)
    }
    ciphertext, err := encrypter.Encrypt(modelJSON)
    if err != nil {
        panic(err)
    }

    fmt.Println(ciphertext)
}


