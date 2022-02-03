package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jaronnie/grpc-gateway-example/config"
	"github.com/jaronnie/grpc-gateway-example/internal"
	"github.com/jaronnie/grpc-gateway-example/pkg/cmd"
)

func main() {
	app := internal.App{}

	if err := cmd.PreBuild(app, &app.Ctx); err != nil {
		fmt.Println(err)
	}

	// config
	cm := config.NewConfigManager(app.Ctx.AppName)
	_ = cm.NewConfig()

	// init app
	if err := app.Init(); err != nil {
		panic(err)
	}

	// run app
	if err := app.Run(); err != nil {
		panic(err)
	}

	signalHandler(&app)
	return

}

func signalHandler(app *internal.App) {
	// signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Println(fmt.Sprintf("service get a signal %s", s.String()))
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			// app exit
			if err := app.Exit(); err != nil {
				panic(err)
			}
			return
		case syscall.SIGHUP:
		// TODO app reload
		default:
			return
		}
	}
}
