package main

import (
	"learnGolang/microservicesWithGo/registration"
	"learnGolang/microservicesWithGo/registration/http/rest"
	"learnGolang/microservicesWithGo/registration/nats"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	notifier := &nats.Notifier{}
	service := registration.NewRegistrationService(notifier)
	regHandler := rest.NewRegistrationHandler(service)
	r := chi.NewRouter()
	r.Post("/", regHandler.ServeHTTP)

	http.ListenAndServe(":8080", r)
}
