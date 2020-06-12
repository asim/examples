module github.com/micro/examples/blog/posts

go 1.13

replace github.com/micro/examples/blog/tags => ../tags

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.0
	github.com/gosimple/slug v1.9.0
	github.com/micro/examples/blog/post v0.0.0-20200611104942-3aa40685d492
	github.com/micro/examples/blog/tags v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.8.0
	google.golang.org/protobuf v1.22.0
)
