package api

import (
	"context"
	"github.com/jaronnie/grpc-gateway-example/internal/proto/pb/corev1"
)

func (c Credential) SayHello(ctx context.Context, request *corev1.HelloRequest) (*corev1.HelloReply, error) {
	return &corev1.HelloReply{Message: request.Name + " world"}, nil
}

func (c Credential) SayHello2(ctx context.Context, request *corev1.HelloRequest) (*corev1.HelloReply, error) {
	return &corev1.HelloReply{Message: request.Name + " world"}, nil
}

func (c Credential) SayHello3(ctx context.Context, request *corev1.HelloRequest) (*corev1.HelloReply, error) {
	return &corev1.HelloReply{Message: request.Name + " world"}, nil
}

func (c Machine) SayHello(ctx context.Context, request *corev1.HelloRequest) (*corev1.HelloReply, error) {
	return &corev1.HelloReply{Message: request.Name + " world"}, nil
}
