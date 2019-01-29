package main

import (
	"fmt"

	hello "github.com/micro/examples/grpc/greeter/srv/proto/hello"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro/metadata"

	"context"
)

func main() {
	service := grpc.NewService()
	service.Init()

	// 使用模板生成的方法存根
	cl := hello.NewSayService("go.micro.srv.greeter", service.Client())

	// 元数据头可以任意设置
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	// 随意声明请求问候的人
	rsp, err := cl.Hello(ctx, &hello.Request{
		Name: "皮皮",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
