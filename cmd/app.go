package cmd

import (
	"context"

	"github.com/paudelgaurav/gin-api-boilerplate/internal"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	internal.Modules,
)

func Execute() {
	internal.LoadEnv()
	app := fx.New(
		fx.Options(
			Modules,
		),
		fx.Invoke(startWebServer),
	)

	app.Run()
}

func startWebServer(lifecycle fx.Lifecycle, server internal.Router) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.RunServer()
				return nil
			},
			OnStop: func(context context.Context) error {
				return nil
			},
		},
	)
}
