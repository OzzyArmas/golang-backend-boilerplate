package main

import (
	"context"

	"github.com/BowlFinder/bowl-finder-server/config"
	"github.com/BowlFinder/bowl-finder-server/db"
	"github.com/BowlFinder/bowl-finder-server/graph"
	"github.com/BowlFinder/bowl-finder-server/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func startEngine(context.Context, *config.Config, *gin.Engine) {
    // zap.Info("Starting program!!!!")
}

func main() {

	app := fx.New(
        fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
            
            return &fxevent.ZapLogger{Logger: log}
        }),
        fx.Provide(
            context.Background, 
            http.Server,
            config.NewConfig,
            db.New,
            zap.NewProduction,
            graph.NewResolver,
        ),
		fx.Invoke(startEngine),
	)
	app.Run()
}
