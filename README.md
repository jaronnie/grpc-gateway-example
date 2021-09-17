# grpc-gateway-example

>  参考 https://grpc-ecosystem.github.io/grpc-gateway/

## 环境

* Go: 1.16
* Protoc: libprotoc 3.16.0
* protoc-gen-go: v1.27.1
* protoc-gen-go-grpc: 1.1.0

## 生成

```shell
protoc -I ./proto \                                          
  --go_out ./proto --go_opt paths=source_relative \
  --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
  ./proto/*.proto
```

