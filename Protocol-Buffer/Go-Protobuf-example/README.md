# Protocol Buffer with Go

## Generate Go Code using Protobuf

    $ brew install protobuf
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    $ go env -w GO111MODULE=off

    $ go mod init Go-Protobuf-example
    $ go mod tidy 

    $ protoc -I src/  --go_out=src src/hello.proto

