package dto

import (
	"database/sql"
)

type Store struct {
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	OpeningHours string  `json:"opening_hours"`
	InsertDate   string
	UpdateDate   string
	DeleteDate   sql.NullString
}
