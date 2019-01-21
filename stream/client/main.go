package main

import (
	"fmt"

	"context"
	proto "examples/stream/server/proto"
	"github.com/micro/go-micro"
)

func bidirectional(cl proto.StreamerService) {
	// create streaming client
	stream, err := cl.Stream(context.Background())
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// bidirectional stream
	// send and receive messages for a 10 count
	for j := 0; j < 10; j++ {
		if err := stream.Send(&proto.Request{Count: int64(j)}); err != nil {
			fmt.Println("err:", err)
			return
		}
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println("recv err", err)
			break
		}
		fmt.Printf("Sent msg %v got msg %v\n", j, rsp.Count)
	}

	// close the stream
	if err := stream.Close(); err != nil {
		fmt.Println("stream close err:", err)
	}
}

func serverStream(cl proto.StreamerService) {
	// send request to stream count of 10
	stream, err := cl.ServerStream(context.Background(), &proto.Request{Count: int64(10)})
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// server side stream
	// receive messages for a 10 count
	for j := 0; j < 10; j++ {
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println("recv err", err)
			break
		}
		fmt.Printf("got msg %v\n", rsp.Count)
	}

	// close the stream
	if err := stream.Close(); err != nil {
		fmt.Println("stream close err:", err)
	}
}

func main() {
	service := micro.NewService()
	service.Init()

	// create client
	cl := proto.NewStreamerService("go.micro.srv.stream", service.Client())

	// bidirectional stream
	bidirectional(cl)

	// server side stream
	serverStream(cl)
}
