package main

import (
	"github.com/MahdiZerara/web-backend/adapter"
	"github.com/MahdiZerara/web-backend/app"
	"github.com/MahdiZerara/web-backend/config"
	"github.com/MahdiZerara/web-backend/controller"
	"github.com/MahdiZerara/web-backend/domain/repository"
	"github.com/MahdiZerara/web-backend/router"
	"github.com/MahdiZerara/web-backend/server"
)

// Handle dependency injection
func buildApp(cfg *config.AppConfig) (*app.App, error) {
	postgresClient := adapter.NewPostgresClient(cfg.PostgresDBConfig)
	if err := postgresClient.Connect(); err != nil {
		return nil, err
	}

	storeRepository := repository.NewStoreRepository(postgresClient)
	storeCtrl := controller.NewStoreController(storeRepository)
	httpCtrl := controller.NewHTTPController(storeCtrl)
	appCtrl := controller.NewAppController(httpCtrl)
	router := router.NewRouter(appCtrl)
	httpServer := server.NewHTTPServer(cfg, router)
	app := app.NewApp(httpServer)
	return app, nil
}
