// Package app package app
package app

import (
	"tpc/pkg/log"
)

var logger = log.New()

func RunApp() {
	logger.Info("hello world")
}
