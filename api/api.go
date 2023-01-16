package api

import (
  "net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)


type (
  IAPIService interface {
    Predict(input string) output string, err error
  }
)

type (
    API struct {
      Router *chi.Mux
      Key string
      service IAPIService
    }
)
func New(r *chi.Mux, key string, srv IAPIService) *API {
  return &API{Router: r, Key: key, service: srv}
} 

func (a *API) Run(port string) error {
  r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/predict", func(w http.ResponseWriter, r *http.Request) {
		// Get input data from the request body
    var input struct{ Input string `json:"input"` }
    if err := r.Bind(&input); err != nil {
      w.Write(err.Error())
      return
    }
    
		// Apply the model to the input data
    output, err := a.Predict(input.Input)
    if err != nil {
      w.Write(err.Error())
      return
    }
    

		// Write the output to the response body
		w.Write(output)
	})

	return http.ListenAndServe(":"+port, r)
}


