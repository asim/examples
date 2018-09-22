# Publisher Service

This is the Publisher service

Generated with

```
micro new github.com/ne0z/go-micro-googlepubsub-demo/Publisher --namespace=go.googlepubsub --alias=publisher --type=srv --plugin=registry=consul:broker=googlepubsub
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.googlepubsub.srv.publisher
- Type: srv
- Alias: publisher

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

Setup environment variables

```console
export PROJECT_ID=your_project_id
export GOOGLE_APPLICATION_CREDENTIALS=~/path/your_project_credentials.json
```

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./publisher-srv
```

Build a docker image
```
make docker
```