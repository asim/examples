module github.com/micro/examples

go 1.13

replace k8s.io/api => k8s.io/api v0.0.0-20190708174958-539a33f6e817

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d

replace k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190708180123-608cd7da68f7

replace k8s.io/client-go => k8s.io/client-go v11.0.0+incompatible

replace k8s.io/component-base => k8s.io/component-base v0.0.0-20190708175518-244289f83105

replace google.golang.org/grpc => google.golang.org/grpc v1.24.0

require (
	github.com/99designs/gqlgen v0.10.1
	github.com/astaxie/beego v1.12.0
	github.com/emicklei/go-restful v2.11.1+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/go-delve/delve v1.4.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/rpc v1.2.0
	github.com/gorilla/websocket v1.4.1
	github.com/gosimple/slug v1.9.0
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711
	github.com/lithammer/shortuuid/v3 v3.0.4 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/examples/blog/post v0.0.0-20200611104942-3aa40685d492 // indirect
	github.com/micro/examples/helloworld v0.0.0-20200611083641-71addf7d37de
	github.com/micro/go-micro/v2 v2.8.0
	github.com/micro/go-plugins/broker/grpc/v2 v2.3.0
	github.com/micro/go-plugins/client/selector/static/v2 v2.3.0
	github.com/micro/go-plugins/config/source/configmap/v2 v2.3.0
	github.com/micro/go-plugins/config/source/grpc/v2 v2.3.0
	github.com/micro/go-plugins/registry/etcd/v2 v2.3.0
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.8.0
	github.com/micro/go-plugins/wrapper/select/roundrobin/v2 v2.3.0
	github.com/micro/go-plugins/wrapper/select/shard/v2 v2.3.0
	github.com/micro/micro/v2 v2.4.0
	github.com/micro/services/helloworld v0.0.0-20200531144117-cc559211129a // indirect
	github.com/micro/services/helloworld/api v0.0.0-20200531144117-cc559211129a // indirect
	github.com/pborman/uuid v1.2.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/vektah/gqlparser v1.2.0
	go.etcd.io/etcd v3.3.22+incompatible // indirect
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.24.0
)

replace github.com/micro/go-micro/v2 => github.com/micro/go-micro/v2 v2.8.1-0.20200604083217-bd3ef6732811
