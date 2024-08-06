package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

func NewServer(lc fx.Lifecycle) *http.Server {
	app := initializeGinServer()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			go app.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown(ctx)
		},
	})

	return app
}

func initializeGinServer() *http.Server {
	app := gin.Default()
	setMiddleware(app)
	app.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	srv := &http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	return srv
}

func setMiddleware(app *gin.Engine) {
	//app.Use(gin.Logger())
}
