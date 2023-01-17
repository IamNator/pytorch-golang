package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)


type (
  IAPIService interface {
    Predict(input string) (output string, err error)
  }
)

type (
    API struct {
      Router *chi.Mux
      service IAPIService
    }
)
func New(r *chi.Mux, srv IAPIService) *API {
  return &API{Router: r, service: srv}
} 

func (a *API) Run(port string) error {

  r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/predict", func(w http.ResponseWriter, r *http.Request) {
		// Get input data from the request body
    var input struct{ Input string `json:"input"` }
    if err := json.NewDecoder(r.Body).Decode(&input); err!= nil {
      w.WriteHeader(http.StatusBadRequest)
      w.Write([]byte(err.Error()))
      return
    }
    
		// Apply the model to the input data
    output, err := a.service.Predict(input.Input)
    if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte(err.Error()))
      return
    }
    

		// Write the output to the response body
		w.WriteHeader(http.StatusOK)
    contentType := "application/json charset=utf-8"
    w.Header().Set("Content-Type", contentType)
    w.Write([]byte(output))
	})

	return http.ListenAndServe(":"+port, r)
}


