package registration

import (
	"log"
	"net/http"
	"strconv"
)

type RegistrationHandler struct{}

func (rh *RegistrationHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Printf("could not parse form because of %v", err)
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	registration := &Registration{}

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
	rw.WriteHeader(http.StatusCreated)
}
