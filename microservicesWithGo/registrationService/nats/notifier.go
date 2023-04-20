package nats

import (
	log "github.com/sirupsen/logrus"
	"learnGolang/microservicesWithGo/registration"

	"github.com/nats-io/nats.go"
)

type Notifier struct{}

func (nn *Notifier) InformAboutNewRegistration(registration *registration.Registration) error {
	registrationLogger := log.WithField("Registration", registration)
	registrationLogger.Info("Inform about new registration.")
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		registrationLogger.WithError(err).Error("Could not connect to server.")
		return err
	}
	defer nc.Close()

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()
	return c.Publish("registration.new", registration)
}
