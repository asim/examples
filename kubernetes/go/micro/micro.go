// Package micro implements a go-micro service for k8s
package micro

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/router/static"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
)

// NewService returns a new go-micro service pre-initialised for k8s
func NewService(opts ...micro.Option) micro.Service {
	// create registry and selector
	r := kubernetes.NewRegistry()

	// set the registry and selector
	options := []micro.Option{
		micro.Registry(r),
		micro.Router(static.NewRouter()),
	}

	// append user options
	options = append(options, opts...)

	// return a micro.Service
	return micro.NewService(options...)
}
