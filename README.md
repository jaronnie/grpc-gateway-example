# grpc-gateway-example

>  参考 https://grpc-ecosystem.github.io/grpc-gateway/

## generate grpc code

```shell
make proto
```

## deploy
```shell
docker build -t jaronnie/kube-grpc-gateway:v2 .
docker run -itd -p 8090:8090 -p 9603:9603 jaronnie/kube-grpc-gateway:v2 
```

## http 调用

```shell
curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": " hello"}'               
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
proto.Core
```

查询服务提供的方法

`grpcurl -plaintext localhost:9603 list proto.Core`

查看更详细的描述

`grpcurl -plaintext localhost:9603 describe proto.Core`

```shell
$ grpcurl -plaintext localhost:9603 describe proto.Core     
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

`grpcurl -plaintext -d '{"name":"hello"}' localhost:9603 proto.Core/SayHello`

```shell
$ grpcurl -plaintext -d '{"name":"hello"}' localhost:9603 proto.Core/SayHello
{
  "message": "hello world"
}
```