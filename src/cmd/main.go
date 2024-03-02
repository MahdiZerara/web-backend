package main

import (
	"context"
	"log"

	"github.com/MahdiZerara/web-backend/config"
)

func main() {
	appCfg, newAppConfigErr := config.NewAppConfig()
	if newAppConfigErr != nil {
		log.Fatalf("An error occurred while retrieving the global configs: %v", newAppConfigErr)
	}

	app, buildAppErr := buildApp(appCfg)
	if buildAppErr != nil {
		log.Fatalf("An error occurred while building the application: %v", buildAppErr)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := app.HTTPServer.Start(ctx); err != nil {
		log.Fatalf(err.Error())
	}
}
