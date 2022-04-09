#!/usr/bin/env bash

protoc -I src/  --go_out=src src/hello.proto

# protoc -I src/  --js_out=src/ src/hello.proto
