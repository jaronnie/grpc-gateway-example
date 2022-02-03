#!/bin/bash

set -e

WorkPath=$(pwd)

# export grpc plugin bin to PATH
PATH="${WorkPath}"/tools/proto/bin:"${PATH}"

# proto tool path
PROTOC_GEN_PATH="$WorkPath"/tools/proto

if [ ! -f "$PROTOC_GEN_PATH"/bin/protoc-gen-go ]; then
  {
    go get github.com/golang/protobuf/protoc-gen-go@v1.3.2 & go build -o "$PROTOC_GEN_PATH"/bin/protoc-gen-go github.com/golang/protobuf/protoc-gen-go
  }
fi

if [ ! -f "$PROTOC_GEN_PATH"/bin/protoc-gen-grpc-gateway ]; then
  {
    go build -o "$PROTOC_GEN_PATH"/bin/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
  }
fi

for pkg in $(find internal -type f | grep "\.proto" | xargs -L 1); do
  {
    echo "generate ==> $pkg"
    cd "${WorkPath}"
    cd "$(dirname "${pkg}")"
    protoc -I. -I/"${WorkPath}"/third/proto/ --go_out=plugins=grpc:. "${pkg##*/}"
    protoc -I. -I/"${WorkPath}"/third/proto/ --grpc-gateway_out=logtostderr=true:. "${pkg##*/}"
  }
done

