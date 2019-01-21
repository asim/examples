package main

import (
	"log"
	"time"

	"context"
	hello "examples/greeter/srv/proto/hello"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/transport"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		// setup a new transport with secure option
		micro.Transport(
			// create new transport
			transport.NewTransport(
				// set to automatically secure
				transport.Secure(true),
			),
		),
	)

	service.Init()

	hello.RegisterSayHandler(service.Server(), new(Say))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
