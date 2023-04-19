package main

import (
	"learnGolang/microservicesWithGo/registration/registration"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	regHandler := &registration.RegistrationHandler{}
	r := chi.NewRouter()
	r.Post("/", regHandler.ServeHTTP)

	http.ListenAndServe(":8080", r)
}
