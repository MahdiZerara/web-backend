package repository

import (
	"github.com/MahdiZerara/web-backend/adapter"
	"github.com/MahdiZerara/web-backend/domain/dto"
)

type StoreRepository interface {
	InsertStore(*dto.Store) error
	RetrieveStores() ([]dto.Store, error)
}

type storeRepository struct {
	DBClient *adapter.PostgresClient
}

// NewStoreRepository creates a new instance of storeRepository.
func NewStoreRepository(dbClient *adapter.PostgresClient) StoreRepository {
	return &storeRepository{
		DBClient: dbClient,
	}
}

func (repo *storeRepository) InsertStore(store *dto.Store) error {
	sqlStatement := `
		INSERT INTO store (name, location, longitude, latitude, opening_hours)
		VALUES ($1, $2, $3, $4, $5)
	`

	return repo.DBClient.Create(
		sqlStatement,
		store.Name,
		store.Location,
		store.Longitude,
		store.Latitude,
		store.OpeningHours,
	)
}

func (repo *storeRepository) RetrieveStores() ([]dto.Store, error) {
	sqlStatement := "SELECT * FROM Store"
	rows, readErr := repo.DBClient.Read(sqlStatement)
	if readErr != nil {
		return nil, readErr
	}
	defer rows.Close()

	stores := make([]dto.Store, 0)
	for rows.Next() {
		var store dto.Store
		if err := rows.Scan(&store.Name, &store.Location, &store.Longitude, &store.Latitude, &store.OpeningHours, &store.InsertDate, &store.UpdateDate, &store.DeleteDate); err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	return stores, nil
}
