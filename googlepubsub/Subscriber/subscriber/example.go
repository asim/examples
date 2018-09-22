package subscriber

import (
	"encoding/json"

	proto "github.com/micro/examples/googlepubsub/Subscriber/proto/example"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/broker"
)

func Handler(p broker.Publication) error {
	request := &proto.Message{}
	err := json.Unmarshal(p.Message().Body, request)
	if err != nil {
		log.Log("Invalid JSON format")
	}
	log.Log("Say: ", request.Say)
	return nil
}
