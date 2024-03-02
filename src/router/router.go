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
	http.HandleFunc("/stores", r.Controller.HTTP.StoreController.HandleRequest)
}
