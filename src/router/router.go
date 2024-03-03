package router

import (
	"net/http"

	"github.com/MahdiZerara/web-backend/controller"
)

type Router struct {
	Controller *controller.AppController
}

func NewRouter(ctrl *controller.AppController) *Router {
	return &Router{
		Controller: ctrl,
	}
}

func (r *Router) InitRoutes() {
	r.initStoreRoutes()
}

func (r *Router) initStoreRoutes() {
	http.HandleFunc("/stores", r.corsMiddleware(r.Controller.HTTP.StoreController.HandleRequest))
}

func (r *Router) corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
