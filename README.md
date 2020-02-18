# Golang fractal implementation

Flactal golang implementation

## generate protobuf

```sh

go get -u github.com/golang/protobuf/protoc-gen-go

protoc -I=./proto --go_out=./proto ./proto/entity.proto

```
