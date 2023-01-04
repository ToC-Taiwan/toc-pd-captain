package main

import (
	"tpc/cmd/config"
	"tpc/internal/app"
)

func main() {
	cfg := config.GetConfig()

	app.InitDB(cfg.Database)
	app.RunApp()
}
