package internal

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(NewDatabase),
	fx.Provide(NewRouter),
)
