package main

import (
	"stories/cmd/config"
	"stories/pkg/stories/server"
)

func serve() {
	cfg := config.NewConfig()
	reporting := newReporting(cfg.GetEnv())
	server.NewServer(cfg, reporting.getLogger()).Start()
}
