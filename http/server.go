package http

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/OzzyArmas/golang-backend-boilerplate/config"
	"github.com/OzzyArmas/golang-backend-boilerplate/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func addGroups(r *gin.Engine, resolver *graph.Resolver) {
	query := r.Group("/query")
	{
		query.POST("/", graphqlHandler(resolver))
	}
	query.GET("/", playgroundHandler())

}

func adminFunction(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"adminFunction": "adminFunction content"})
}
func usersFunction(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"usersFunction": "usersFunction content"})
}

func graphqlHandler(resolver *graph.Resolver) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func Server(
	ctx context.Context,
	lc fx.Lifecycle,
	cfg *config.Config,
	log *zap.Logger,
	resolver *graph.Resolver,
) *gin.Engine {
	router := gin.Default()
	addGroups(router, resolver) // define rules for router
	log.Info(cfg.HTTP.ListenPort)
	srv := &http.Server{Addr: cfg.HTTP.ListenPort, Handler: router} // define a web server

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr) // the web server starts listening on 8080
			if err != nil {
				fmt.Println("[My Demo] Failed to start HTTP Server at", srv.Addr)
				return err
			}
			go srv.Serve(ln) // process an incoming request in a go routine
			fmt.Println("[My Demo]Succeeded to start HTTP Server at", srv.Addr)
			return nil

		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(ctx) // stop the web server
			fmt.Println("[My Demo] HTTP Server is stopped")
			return nil
		},
	})

	return router
}
