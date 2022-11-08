package internal

import (
	"context"
	"github.com/jaronnie/grpc-gateway-example/internal/proto/pb/corev1"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jaronnie/grpc-gateway-example/internal/api"
	"github.com/jaronnie/grpc-gateway-example/pkg/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (app *App) grpcServer() (s *grpc.Server, err error) {
	logx.Infof("start grpc server 0.0.0.0:%s\n", app.GrpcPort)
	listen, err := net.Listen("tcp", "0.0.0.0:"+app.GrpcPort)
	if err != nil {
		panic(err)
	}

	s = grpc.NewServer()
	// Register reflection service on gRPC server.
	reflection.Register(s)

	machine := &api.Machine{}
	corev1.RegisterMachineServer(s, machine)

	credential := &api.Credential{}
	corev1.RegisterCredentialServer(s, credential)

	app.GrpcServer = s

	return s, s.Serve(listen)
}

func (app *App) gatewayServer() (s *http.Server, err error) {
	logx.Infof("start gateway server 0.0.0.0:%s\n", app.HttpPort)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := corev1.RegisterMachineHandlerFromEndpoint(context.Background(), mux, "0.0.0.0:"+app.GrpcPort, opts); err != nil {
		return nil, err
	}

	if err := corev1.RegisterCredentialHandlerFromEndpoint(context.Background(), mux, "0.0.0.0:"+app.GrpcPort, opts); err != nil {
		return nil, err
	}

	s = &http.Server{
		Addr:    "0.0.0.0:" + app.HttpPort,
		Handler: mux,
	}

	app.HTTPServer = s

	return s, s.ListenAndServe()
}
