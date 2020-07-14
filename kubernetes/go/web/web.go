package web

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/router/static"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
)

// NewService returns a web service for kubernetes
func NewService(opts ...web.Option) web.Service {
	// setup
	k := kubernetes.NewRegistry()

	// create new service
	service := micro.NewService(
		micro.Registry(k),
		micro.Router(static.NewRouter()),
	)

	// prepend option
	options := []web.Option{
		web.MicroService(service),
	}

	options = append(options, opts...)

	// return new service
	return web.NewService(options...)
}
