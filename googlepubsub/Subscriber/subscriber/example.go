package subscriber

import (
	"encoding/json"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/broker"
	proto "github.com/ne0z/examples/googlepubsub/Subscriber/proto/example"
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
