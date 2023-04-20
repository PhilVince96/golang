package rest

import (
	"learnGolang/microservicesWithGo/registration"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type RegistrationHandler struct {
	service *registration.RegistrationService
}

func NewRegistrationHandler(service *registration.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{
		service,
	}
}

func (rh *RegistrationHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Printf("could not parse form because of %v", err)
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	registration := &registration.Registration{}

	// read from form
	registration.Firstname = req.Form.Get("Firstname")
	registration.Lastname = req.Form.Get("Lastname")
	registration.Email = req.Form.Get("Email")
	registration.Company = req.Form.Get("Company")
	registration.Date = req.Form.Get("Date")
	registration.Trainingcode = req.Form.Get("Trainingcode")

	b, err := strconv.ParseBool(req.Form.Get("PrivacyProtectionAccepted"))
	if err != nil {
		log.Printf("Could not parse value for PrivacyProtectionAccepted because of %v", err)
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	registration.PrivacyProctectionAccepted = b
	log.Printf("new registration %+v", registration)

	err = rh.service.HandleNewRegistration(registration)
	if err != nil {
		log.Printf("Could not handle registration: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}
