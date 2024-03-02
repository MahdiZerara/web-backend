package controller

type AppController struct {
	HTTP *HTTPController
}

func NewAppController(httpCtrl *HTTPController) *AppController {
	return &AppController{
		HTTP: httpCtrl,
	}
}

type HTTPController struct {
	StoreController *StoreController
}

func NewHTTPController(storeCtrl *StoreController) *HTTPController {
	return &HTTPController{
		StoreController: storeCtrl,
	}
}
