package main

import (
	"learnGolang/microservicesWithGo/registration"
	"learnGolang/microservicesWithGo/registration/http/rest"
	"learnGolang/microservicesWithGo/registration/nats"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	notifier := &nats.Notifier{}
	service := registration.NewRegistrationService(notifier)
	regHandler := rest.NewRegistrationHandler(service)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.AllowContentType("application/x-www-form-urlencoded"))
	r.Post("/", regHandler.ServeHTTP)

	http.ListenAndServe(":8080", r)
}
