# gRPC 网关

本目录是gRPC网关示例，该示例通过[grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)生成。

服务使用[micro/go-grpc](https://github.com/micro/go-grpc)编写，与gRPC-gateway及其它gRPC服务兼容。
                              
请查看[grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)了解如何生成网关。这里我们已经使用greeter服
务的proto原型文件生成了网关代码，不过在网关中适当加一些选项。

网关的proto

```proto
service Say {
	rpc Hello(Request) returns (Response) {
		option (google.api.http) = {
			post: "/greeter/hello"
			body: "*"
		};
	}
}
```

greeter的proto

```proto
service Say {
	rpc Hello(Request) returns (Response) {}
}
```

可以查看两个目录的`hello.proto`了解，gateway的多的选项是用来告知插件生成代码时`Hello`处理器会被映射为`/greeter/hello`路由。

## 使用

运行`go.micro.srv.greeter`服务

```bash
go run ../greeter/srv/main.go --server_address=localhost:9090
```

运行网关

```bash
go run main.go
```

我们指定网关的服务地址为`localhost:8080`，详见代码[main.go](main.go)文件

通过网关请求greeter服务

```bash
curl -d '{"name": "john"}' http://localhost:8080/greeter/hello
```


