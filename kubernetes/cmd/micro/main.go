package main

import (
	"os"

	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
	cli "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/server"
	srv "github.com/micro/go-micro/v2/server/grpc"
	bkr "github.com/micro/go-plugins/broker/grpc/v2"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"

	// disable namespace by default
	_ "github.com/micro/go-micro/v2/api"
)

func main() {

	// set values for registry/selector
	os.Setenv("MICRO_REGISTRY", "kubernetes")
	os.Setenv("MICRO_ROUTER", "static")

	// setup broker/client/server
	broker.DefaultBroker = bkr.NewBroker()
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()

	// init command
	cmd.Init()
}
