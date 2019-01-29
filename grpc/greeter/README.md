# Greeter 问候服务

基于gRPC的Go-Micro服务示例

## 内容

- **srv** - gRPC 问候服务 服务端
- **cli** - gRPC 调用问候服务的客户端（只调用一次）

### 运行服务

运行服务端

```
$ go run srv/main.go --registry=mdns
2016/11/03 18:41:22 Listening on [::]:55194
2016/11/03 18:41:22 Broker Listening on [::]:55195
2016/11/03 18:41:22 Registering node: go.micro.srv.greeter-1e200612-a1f5-11e6-8e84-68a86d0d36b6
```

运行命令中的`--registry=mdns`flag指定注册中心为组播模式，相当于局域网模式的服务发现。

### 测试服务

```
$ go run cli/main.go --registry=mdns
Hello John
```

