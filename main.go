package main

import (
	"github.com/OzzyArmas/golang-backend-boilerplate/config"
	"github.com/OzzyArmas/golang-backend-boilerplate/db"
	"github.com/OzzyArmas/golang-backend-boilerplate/graph"
	"github.com/OzzyArmas/golang-backend-boilerplate/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {

	app := fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {

			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			http.Server,
			config.NewConfig,
			db.New,
			zap.NewProduction,
			graph.NewResolver,
		),
		fx.Invoke(func(config.Config, *gin.Engine) {}),
	)
	app.Run()
}
