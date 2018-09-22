package main

import (
	"encoding/json"
	"os"

	"github.com/go-log/log"
	proto "github.com/micro/examples/googlepubsub/Publisher/proto/example"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/googlepubsub"
)

func main() {

	projectID := googlepubsub.ProjectID(os.Getenv("PROJECT_ID"))
	gpubsub := googlepubsub.NewBroker(projectID)

	msgStuct := &proto.Message{
		Say: "hello",
	}
	msg, _ := json.Marshal(msgStuct)
	log.Log(string(msg))
	pub := &broker.Message{
		Header: map[string]string{
			"id": "1",
		},
		Body: msg,
	}
	gpubsub.Publish("go.googlepubsub.srv.subscriber", pub)
}
