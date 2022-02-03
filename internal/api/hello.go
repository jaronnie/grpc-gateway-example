package api

import (
	"context"

	"github.com/jaronnie/grpc-gateway-example/internal/proto"
)

func (c Core) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: request.Name + " world"}, nil
}
