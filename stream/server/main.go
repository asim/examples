package main

import (
	"log"

	"context"
	"github.com/micro/go-micro"

	proto "examples/stream/server/proto"
)

type Streamer struct{}

// Server side stream
func (e *Streamer) ServerStream(ctx context.Context, req *proto.Request, stream proto.Streamer_ServerStreamStream) error {
	log.Printf("Got msg %v", req.Count)
	for i := 0; i < int(req.Count); i++ {
		if err := stream.Send(&proto.Response{Count: int64(i)}); err != nil {
			return err
		}
	}
	return nil
}

// Bidirectional stream
func (e *Streamer) Stream(ctx context.Context, stream proto.Streamer_StreamStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Got msg %v", req.Count)
		if err := stream.Send(&proto.Response{Count: req.Count}); err != nil {
			return err
		}
	}
}

func main() {
	// new service
	service := micro.NewService(
		micro.Name("go.micro.srv.stream"),
	)

	// Init command line
	service.Init()

	// Register Handler
	proto.RegisterStreamerHandler(service.Server(), new(Streamer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
