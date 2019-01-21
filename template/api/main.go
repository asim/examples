package main

import (
	"github.com/micro/go-log"

	"examples/template/api/client"
	"examples/template/api/handler"
	"github.com/micro/go-micro"

	example "examples/template/api/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.template"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(client.ExampleWrapper(service)),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
