# Protocol Buffer with Go

## Generate Go Code using Protobuf

    $ pwd
    /Users/Shiru/go/src/github.com/Shiru99/Go-Protobuf-example

    $ brew install protobuf
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    $ go env -w GO111MODULE=off

    $ protoc -I src/  --go_out=src src/hello.proto
    $ go get -u google.golang.org/protobuf/runtime/protoimpl 
    $ go get -u google.golang.org/protobuf/reflect/protoreflect

