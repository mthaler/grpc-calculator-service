# grpc-calculator-service
gRPC calculator service from gRPC Master Class: Build Modern API &amp; Microservices on Udemy 

Install the protobuf compiler plugins for Go using the following commands:

```bash
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Update PATH so that the protoc compiler can find the plugins:

```bash
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Compile a gRPC service

```bash
$ protoc --go_out=. --go-grpc_out=.  calculatorpb/calculator.proto
```
## Reflection

gRPC supports reflection to query available APIs.

### Evans

[Evans](https://github.com/ktr0731/evans/) is a CLI that can be used to query gRPC APIs.
