package main

import (
	"net/http"

	"github.com/IamNator/pytorch-golang/api"
	"github.com/IamNator/pytorch-golang/service"
	"github.com/go-chi/chi"
)


func main() {

	newService := service.New()

	r := chi.NewRouter()

	newAPI := api.New(r, newService)

	log.Fatal(newAPI.Run("3000"))
}

