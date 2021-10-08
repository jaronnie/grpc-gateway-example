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

## http 调用

```shell
curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": " hello"}'               

curl -X POST -k http://localhost:8090/v1/init/crt -d '{"id":1, "name":"11", "type":"input"}'
```
## grpc 调用

> 参考：https://blog.csdn.net/m0_37556444/article/details/100064623

使用 grpcurl 测试代码

安装 grpcurl
```shell
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

查看服务列表

`grpcurl -plaintext localhost:9603 list`

```shell
$ grpcurl -plaintext localhost:9603 list 
grpc.reflection.v1alpha.ServerReflection
proto.Greeter
```

查询服务提供的方法

`grpcurl -plaintext localhost:9603 list proto.Greeter`

查看更详细的描述

`grpcurl -plaintext localhost:9603 describe proto.Greeter`

```shell
$ grpcurl -plaintext localhost:9603 describe proto.Greeter     
proto.Greeter is a service:
service Greeter {
  rpc InitCredential ( .proto.InitCredentialRequest ) returns ( .proto.InitCredentialReply ) {
    option (.google.api.http) = { post:"/v1/init/crt" body:"*"  };
  }
  rpc SayHello ( .proto.HelloRequest ) returns ( .proto.HelloReply ) {
    option (.google.api.http) = { post:"/v1/example/echo" body:"*"  };
  }
}
```

获取类型信息

`grpcurl -plaintext localhost:9603 describe proto.HelloRequest`

```shell
$ grpcurl -plaintext localhost:9603 describe proto.HelloRequest
proto.HelloRequest is a message:
message HelloRequest {
  string name = 1;
}
```

调用服务方法

`grpcurl -plaintext -d '{"name":"hello"}' localhost:9603 proto.Greeter/SayHello`

```shell
$ grpcurl -plaintext -d '{"name":"hello"}' localhost:9603 proto.Greeter/SayHello
{
  "message": "hello world"
}
```

