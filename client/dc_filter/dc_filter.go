package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/router"
	"github.com/micro/go-micro/v2/selector"

	example "github.com/micro/examples/server/proto/example"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// A Wrapper that creates a Datacenter Selector Option
type dcWrapper struct {
	client.Client
}

func (dc *dcWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)

	filter := func(routes []router.Route) []router.Route {
		var result []router.Route

		for _, r := range routes {
			if r.Metadata["datacenter"] == md["datacenter"] {
				result = append(result, r)
			}
		}

		return result
	}

	callOptions := append(opts, client.WithSelectOptions(selector.WithFilter(filter)))

	fmt.Printf("[DC Wrapper] filtering for datacenter %s\n", md["datacenter"])
	return dc.Client.Call(ctx, req, rsp, callOptions...)
}

func NewDCWrapper(c client.Client) client.Client {
	return &dcWrapper{c}
}

func call(i int) {
	// Create new request to service go.micro.srv.example, method Example.Call
	req := client.NewRequest("go.micro.srv.example", "Example.Call", &example.Request{
		Name: "John",
	})

	// create context with metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"datacenter": "local",
	})

	rsp := &example.Response{}

	// Call service
	if err := client.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("Call:", i, "rsp:", rsp.Msg)
}

func main() {
	cmd.Init()

	client.DefaultClient = client.NewClient(
		client.Wrap(NewDCWrapper),
	)

	fmt.Println("\n--- Call example ---")
	for i := 0; i < 10; i++ {
		call(i)
	}
}
