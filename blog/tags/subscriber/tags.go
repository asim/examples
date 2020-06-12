package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	tags "github.com/micro/examples/blog/tags/proto/tags"
)

type Tags struct{}

func (e *Tags) Handle(ctx context.Context, msg *tags.IncreaseCountRequest) error {
	log.Info("Handler Received message")
	return nil
}

func Handler(ctx context.Context, msg *tags.IncreaseCountRequest) error {
	log.Info("Function Received message: ")
	return nil
}
