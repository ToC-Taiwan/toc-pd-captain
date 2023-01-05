// Package app package app
package app

import (
	"os"
	"os/signal"
	"syscall"

	"tpc/cmd/config"
	"tpc/internal/controller/http/router"
	"tpc/internal/usecase/user"
	"tpc/pkg/httpserver"
	"tpc/pkg/log"

	"github.com/gin-gonic/gin"
)

var logger = log.Get()

func RunApp(cfg *config.Config) {
	gin.SetMode(cfg.Server.RouterDebugMode)

	// create all needed usecases
	userUseCase := user.NewUserUseCase()

	handler := gin.New()
	router.NewRouter(handler, userUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Server.HTTP))
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
