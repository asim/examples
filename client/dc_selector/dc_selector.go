package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/router"
	"github.com/micro/go-micro/v2/selector"

	example "github.com/micro/examples/server/proto/example"
)

// Built in random hashed node selector
type dcSelector struct {
	opts selector.Options
}

var (
	datacenter = "local"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func (n *dcSelector) Init(opts ...selector.Option) error {
	for _, o := range opts {
		o(&n.opts)
	}
	return nil
}

func (n *dcSelector) Options() selector.Options {
	return n.opts
}

func (n *dcSelector) Select(routes []router.Route, opts ...selector.SelectOption) (*router.Route, error) {
	// Filter the nodes for datacenter {
	for _, r := range routes {
		if r.Metadata["datacenter"] == datacenter {
			return &r, nil
		}
	}

	return nil, selector.ErrNoneAvailable
}

func (n *dcSelector) Record(router.Route, error) error {
	return nil
}

func (n *dcSelector) Close() error {
	return nil
}

func (n *dcSelector) String() string {
	return "dc"
}

// Return a new selector
func DCSelector(opts ...selector.Option) selector.Selector {
	var sopts selector.Options
	for _, opt := range opts {
		opt(&sopts)
	}
	return &dcSelector{sopts}
}

func call(i int) {
	// Create new request to service go.micro.srv.example, method Example.Call
	req := client.NewRequest("go.micro.srv.example", "Example.Call", &example.Request{
		Name: "John",
	})

	rsp := &example.Response{}

	// Call service
	if err := client.Call(context.Background(), req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("Call:", i, "rsp:", rsp.Msg)
}

func main() {
	cmd.Init()

	client.DefaultClient = client.NewClient(
		client.Selector(DCSelector()),
	)

	fmt.Println("\n--- Call example ---")
	for i := 0; i < 10; i++ {
		call(i)
	}
}
