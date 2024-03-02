package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/MahdiZerara/web-backend/config"
	"github.com/MahdiZerara/web-backend/router"
)

type HTTPServer struct {
	cfg    *config.AppServerConfig
	router *router.Router
}

func NewHTTPServer(
	cfg *config.AppConfig,
	r *router.Router,
) *HTTPServer {
	return &HTTPServer{
		cfg:    cfg.AppServerConfig,
		router: r,
	}
}

func (s *HTTPServer) Start(ctx context.Context) error {
	httpPort := s.cfg.GetRESTPort()
	s.router.InitRoutes()

	log.Printf("Server is running on http://localhost:%s", httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil); err != nil {
		return fmt.Errorf("an error occurred while starting the application HTTP server: %v", err)
	}

	return nil
}
