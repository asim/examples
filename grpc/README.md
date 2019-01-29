# GRPC

该目录下包含使用[go-grpc](https://github.com/micro/go-grpc)的例子。

- [greeter](greeter) - 问候程序示例
- [gateway](gateway) - gRPC网关示例
- [sidecar](sidecar) -【未完成】通过micro sidecar（挎斗，一种设计模式）来提供HTTP服务

## 新建服务

签出[greeter](greeter)示例，里面使用了go-grpc

### 导入 go-grpc

```
import "github.com/micro/go-grpc"
```

### 创建micro.Service服务

```
service := grpc.NewService()
```

## 已有服务

如果要在已有服务中增加grpc该怎么弄？使用plugin的[build构建模式](https://github.com/micro/examples/tree/cn-lang/plugins/#build-pattern)，
但是要把client/server换掉，下面的例子只是为了演示如何声明客户端与服务端，真实情况可能是二者皆有或中有其一，且一般不写在plugins.go里。

### 创建插件文件

创建插件文件plugins.go，它主要用来引入插件

```
package main

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	cli "github.com/micro/go-plugins/client/grpc"
	srv "github.com/micro/go-plugins/server/grpc"
)

func init() {
	// 设置默认的客户端
	client.DefaultClient = cli.NewClient()
	// 设置默认的服务端
	server.DefaultServer = srv.NewServer()
}
```

### 编译成二进制

```
// 本地运行编译
go build -i -o service ./main.go ./plugins.go
```

### 运行

因为client/server已经被换掉了，所以可以直接像平常应用一样执行

```
./service
```
