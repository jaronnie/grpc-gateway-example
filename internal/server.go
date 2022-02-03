package internal

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jaronnie/grpc-gateway-example/internal/api"
	"github.com/jaronnie/grpc-gateway-example/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (app *App) grpcServer() (s *grpc.Server, err error) {
	fmt.Println("start grpc server 0.0.0.0:", app.GrpcPort)
	listen, err := net.Listen("tcp", "0.0.0.0:"+app.GrpcPort)
	if err != nil {
		panic(err)
	}

	s = grpc.NewServer()
	// Register reflection service on gRPC server.
	reflection.Register(s)

	core := &api.Core{}
	proto.RegisterCoreServer(s, core)

	app.GrpcServer = s

	return s, s.Serve(listen)
}

func (app *App) gatewayServer() (s *http.Server, err error) {
	fmt.Println("start gateway server 0.0.0.0:", app.HttpPort)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := proto.RegisterCoreHandlerFromEndpoint(context.Background(), mux, "0.0.0.0:"+app.GrpcPort, opts); err != nil {
		return nil, err
	}

	s = &http.Server{
		Addr:    "0.0.0.0:" + app.HttpPort,
		Handler: mux,
	}

	app.HTTPServer = s

	return s, s.ListenAndServe()
}
