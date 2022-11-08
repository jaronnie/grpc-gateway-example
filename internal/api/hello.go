package api

import (
	"context"
	"github.com/jaronnie/grpc-gateway-example/internal/proto/pb/corev1"
)

func (c Corev1) SayHello(ctx context.Context, request *corev1.HelloRequest) (*corev1.HelloReply, error) {
	return &corev1.HelloReply{Message: request.Name + " world"}, nil
}

func (c Corev1) SayHello2(ctx context.Context, request *corev1.HelloRequest) (*corev1.HelloReply, error) {
	return &corev1.HelloReply{Message: request.Name + " world"}, nil
}

func (c Corev1) SayHello3(ctx context.Context, request *corev1.HelloRequest) (*corev1.HelloReply, error) {
	return &corev1.HelloReply{Message: request.Name + " world"}, nil
}
