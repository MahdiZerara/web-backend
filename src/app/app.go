package app

import (
	"github.com/MahdiZerara/web-backend/server"
)

type App struct {
	HTTPServer *server.HTTPServer
}

func NewApp(httpServer *server.HTTPServer) *App {
	return &App{
		HTTPServer: httpServer,
	}
}
