package main

import (
	"html/template"
	"learnGolang/microservicesWithGo/registration"
	"learnGolang/microservicesWithGo/registration/http/rest"
	"learnGolang/microservicesWithGo/registration/nats"
	"net/http"

	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	notifier := &nats.Notifier{}
	service := registration.NewRegistrationService(notifier)
	regHandler := rest.NewRegistrationHandler(service)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(rest.RateLimiter)
	r.Use(chiprometheus.NewMiddleware("serviceName"))
	// r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.AllowContentType("application/x-www-form-urlencoded"))
	r.Post("/", regHandler.ServeHTTP)
	r.Get("/registration", func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("templates/registration.gohtml")
		if err != nil {
			log.Errorf("Could not parse template files. Error: %v", err)
		}
		template.Execute(w, nil)
	})
	r.Get("/metrics", promhttp.Handler().ServeHTTP)

	http.ListenAndServe(":8080", r)
}
