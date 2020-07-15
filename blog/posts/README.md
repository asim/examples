# Post Service

## Create a post

```
micro call go.micro.service.posts Posts.Post '{"post":{"id":"1","title":"How to Micro","content":"Simply put, Micro is awesome."}}'
micro call go.micro.service.posts Posts.Post '{"post":{"id":"2","title":"Fresh posts are fresh","content":"This post is fresher than the How to Micro one"}}'
```

## Query posts

```
micro call go.micro.service.posts Posts.Query '{}'
micro call go.micro.service.posts Posts.Query '{"slug":"how-to-micro"}'
micro call go.micro.service.posts Posts.Query '{"offset": 10, "limit": 10}'
```

## Delete posts

```
micro call go.micro.service.posts Posts.Delete '{"offset": 10, "limit": 10}'
```

Generated with

```
micro new --namespace=go.micro --type=service posts
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.posts
- Type: service
- Alias: posts

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./posts-service
```

Build a docker image
```
make docker
```
