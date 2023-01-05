// Package app package app
package app

import (
	"os"
	"os/signal"
	"syscall"

	"tpc/cmd/config"
	v1 "tpc/internal/controller/http/v1"
	"tpc/pkg/httpserver"
	"tpc/pkg/log"

	"github.com/gin-gonic/gin"
)

var logger = log.New()

func RunApp(cfg *config.Config) {
	gin.SetMode(cfg.Server.RouterDebugMode)

	handler := gin.New()
	v1.NewRouter(handler)
	httpServer := httpserver.New(
		handler,
		httpserver.Port(cfg.Server.HTTP),
	)

	Serve(httpServer)
}

func Serve(httpServer *httpserver.Server) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info(s.String())
	case err := <-httpServer.Notify():
		logger.Error(err)
	}

	if err := httpServer.Shutdown(); err != nil {
		logger.Error(err)
	}
}
