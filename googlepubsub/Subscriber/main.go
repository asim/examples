package main

import (
	"os"

	"github.com/ne0z/examples/googlepubsub/Subscriber/subscriber"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/googlepubsub"
)

func main() {

	projectID := googlepubsub.ProjectID(os.Getenv("PROJECT_ID"))
	gpubsub := googlepubsub.NewBroker(projectID)

	var sub broker.Subscriber
	// New Service
	service := micro.NewService(
		micro.Name("go.googlepubsub.srv.subscriber"),
		micro.Version("latest"),
		micro.Broker(gpubsub),
		micro.BeforeStop(func() error {
			err := sub.Unsubscribe()
			if err != nil {
				log.Fatal(err)
				return err
			}
			log.Log("Unsubscribe Google Pub/Sub : ", sub.Topic())
			return nil
		}),
	)
	service.Init()

	sub, err := gpubsub.Subscribe("go.googlepubsub.srv.subscriber", subscriber.Handler)
	if err != nil {
		log.Fatal(err)
	}
	log.Log("Subscribe Google Pub/Sub : ", sub.Topic())

	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
