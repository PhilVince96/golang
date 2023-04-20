package nats

import (
	"learnGolang/microservicesWithGo/registration"
	"log"

	"github.com/nats-io/nats.go"
)

type Notifier struct{}

func (nn *Notifier) InformAboutNewRegistration(registration *registration.Registration) error {
	log.Println("Inform about new registration")
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("Could not connect to server: ", err)
		return err
	}
	defer nc.Close()

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()
	return c.Publish("registration.new", registration)
}
