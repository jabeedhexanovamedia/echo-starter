package main

import (
	"github.com/jabeedhexanovamedia/echo-starter/internal/config"
	"github.com/jabeedhexanovamedia/echo-starter/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	e := server.New(cfg)
	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
