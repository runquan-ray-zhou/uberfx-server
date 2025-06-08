package main

import (
	"github.com/runquan-ray-zhou/uberfx-server/httphandler/api/rest"
	"github.com/runquan-ray-zhou/uberfx-server/httpserver"
	"go.uber.org/fx"
)

func main() {
	fx.New(opts()).Run()
}

func opts() fx.Option {
	return fx.Options(
		rest.Module,
		httpserver.Module,
	)
}
