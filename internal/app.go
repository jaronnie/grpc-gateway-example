package internal

import (
	"net/http"

	"github.com/jaronnie/grpc-gateway-example/pkg/logx"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Application base app framework interface
type Application interface {
	Init() error
	Run() error
	Exit() error
}

// App is main service
type App struct {
	Ctx        Context
	GrpcServer *grpc.Server
	HTTPServer *http.Server
	GrpcPort   string
	HttpPort   string
}

// Context is main application context
type Context struct {
	AppName string

	// Magic context for hack
	Magic map[string]interface{}
}

// Init application init
func (app *App) Init() (err error) {
	logx.Logger()

	logx.Infof("app init")
	app.GrpcPort = viper.GetString("grpc.port")
	app.HttpPort = viper.GetString("http.port")
	return
}

// Run application run
func (app *App) Run() (err error) {
	logx.Infof("app run")
	go func() {
		_, err := app.grpcServer()
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		_, err := app.gatewayServer()
		if err != nil {
			panic(err)
		}
	}()
	return nil
}

// Exit application exit
func (app *App) Exit() (err error) {
	return nil
}
