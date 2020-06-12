package main

import (
	"github.com/micro/examples/blog/web/handler"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.blog"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	handl := handler.Handler{
		Client: service.Options().Service.Client(),
	}
	// register call handler
	service.HandleFunc("/post/", handl.Post)
	service.HandleFunc("/new-post", handl.NewPost)
	service.HandleFunc("/edit/", handl.EditPost)
	//service.HandleFunc("/edit-post", handl.EditPost)
	service.HandleFunc("/api/post", handl.PostAPI)
	service.HandleFunc("/", handl.Index)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
