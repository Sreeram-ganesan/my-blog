package apiserver

import (
	"context"
	"fmt"

	"github.com/Sreeram-ganesan/my-blog/internal/adapters/apiserver/internal"
	"github.com/Sreeram-ganesan/my-blog/internal/core/di"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Start(_ context.Context, di *di.DI) {

	r := gin.Default()
	r.Use(zapLoggerMiddleware(zap.S()))
	r.Use(gin.Recovery())

	apiRoute(r, di)

	func() {
		r.Use(static.Serve("/", static.LocalFile("./web/dist", true)))
		r.NoRoute(func(c *gin.Context) {
			c.File("./web/dist/index.html")
		})
	}() // file server handler to serve web application

	listenAddr := fmt.Sprintf("%s:%d", di.Config.Server.Addr, di.Config.Server.Port)
	srv := NewHttpServer(listenAddr, r)
	srv.ShutdownCallback = func() {
		zap.S().Info("Cleaning up resources")
		di.Close()
		zap.S().Infof("Resources has been cleaned up")
	}
	zap.S().Info("Starting HTTP server...")
	go func() {
		srv.start()
	}()

	srv.waitWithGracefulShutdown()
}

func apiRoute(r *gin.Engine, di *di.DI) {
	r.GET("/api/version", internal.GetVersion())
	r.Group("/api/contacts").
		POST("", internal.CreateContact(di.UseCases)).
		GET("", internal.ListAllContacts(di.UseCases)).
		PUT("/:id", internal.UpdateContact(di.UseCases)).
		GET("/:id", internal.GetContact(di.UseCases)).
		DELETE("/:id", internal.DeleteContact(di.UseCases))
	r.Group("/api/blogs").POST("", internal.CreateBlog(di.UseCases))
}
