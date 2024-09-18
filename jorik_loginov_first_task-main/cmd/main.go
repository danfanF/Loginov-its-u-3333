package main

import (
	httpserver "github.com/akayo16/jorik_loginov_first_task/cmd/http"
	postgresqlstorage "github.com/akayo16/jorik_loginov_first_task/cmd/storage/postgresql"
	"github.com/akayo16/jorik_loginov_first_task/config"
	_ "github.com/akayo16/jorik_loginov_first_task/docs/transfer/http/v1/swagger"
	"github.com/akayo16/jorik_loginov_first_task/pkg/lib/logger"
)

// @title Test Example API
// @version 1.0
// @description This Is A Driving Style Design Server.
// @BasePath /api/v1
func main() {
	cfg := config.MustLoad()

	logger.Init(cfg.DevelopmentEnvironment)

	postgresqlstorage.InitDatabase(cfg.Postgresql)

	httpserver.InitHttpServer(cfg.HTTPServer.Port)
}
