package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MahdiZerara/web-backend/domain/dto"
	"github.com/MahdiZerara/web-backend/domain/repository"
)

// StoreController handles HTTP requests for the store entity.
type StoreController struct {
	Repository repository.StoreRepository
}

// NewStoreController creates a new UserController.
func NewStoreController(r repository.StoreRepository) *StoreController {
	return &StoreController{
		Repository: r,
	}
}

func (ctrl *StoreController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// query := r.URL.Query()
		// from := query.Get("from")
		// to := query.Get("to")

		stores, retrieveStoresErr := ctrl.Repository.RetrieveStores(5, 10)
		if retrieveStoresErr != nil {
			fmt.Errorf("error retrieving store %v", retrieveStoresErr)

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{
				"success":false
			}`))
			return
		}

		err := json.NewEncoder(w).Encode(stores)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{
				"success":false
			}`))
			return
		}

	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{
				"success":false
			}`))
			return
		}
		defer r.Body.Close()

		// Unmarshal the JSON into a struct
		var payload dto.Store
		err = json.Unmarshal(body, &payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{
				"success":false
			}`))
			return
		}

		InsertStoreErr := ctrl.Repository.InsertStore(&dto.Store{
			Name:         payload.Name,
			Location:     payload.Location,
			Longitude:    payload.Longitude,
			Latitude:     payload.Latitude,
			OpeningHours: payload.OpeningHours,
		})
		if InsertStoreErr != nil {
			fmt.Errorf("error retrieving store %v", InsertStoreErr)

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{
				"success":false
			}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"success":true
		}`))
	default:
		http.Error(w, "Unsupported request method.", http.StatusMethodNotAllowed)
	}

	return
}
