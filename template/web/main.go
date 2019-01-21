package main

import (
	"github.com/micro/go-log"
	"net/http"

	"examples/template/web/handler"
	"github.com/micro/go-web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.template"),
		web.Version("latest"),
	)

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/example/call", handler.ExampleCall)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
