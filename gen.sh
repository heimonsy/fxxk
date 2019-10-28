#!/bin/sh

# shellcheck disable=SC2128
cd "$(dirname "${BASH_SOURCE}")"

protoc --plugin=protoc-gen-grpc-go="$GOPATH/bin/protoc-gen-go" --proto_path . --go_out=plugins=grpc:"$GOPATH/src" fxxk.proto
