package main

import (
	"github.com/runquan-ray-zhou/uberfx-server/httphandler/api/linknyc"
	"github.com/runquan-ray-zhou/uberfx-server/httphandler/api/pocketdictionary"
	"github.com/runquan-ray-zhou/uberfx-server/httphandler/api/quizme"
	"github.com/runquan-ray-zhou/uberfx-server/httphandler/api/rest"
	"github.com/runquan-ray-zhou/uberfx-server/httphandler/api/rrunquanzhou"
	"github.com/runquan-ray-zhou/uberfx-server/httpserver"
	"github.com/runquan-ray-zhou/uberfx-server/util/cron"
	"go.uber.org/fx"
)

func main() {
	fx.New(opts()).Run()
}

func opts() fx.Option {
	return fx.Options(
		linknyc.Module,          // linknyc
		pocketdictionary.Module, // pocket dictionary
		quizme.Module,           // quiz-me
		rrunquanzhou.Module,     // personal website
		rest.Module,             // health check
		httpserver.Module,       // httpserver
		cron.Module,             // cron job
	)
}
