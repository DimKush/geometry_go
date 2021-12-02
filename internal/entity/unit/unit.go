package unit

import (
	"encoding/json"
)

type Unit struct {
	Id       int              `json:"id"`
	Name     string           `json:"name"`
	GeomJson *json.RawMessage `json:"zone" gorm:"-"`
	Geom     string           `json:"-" gorm:"geom"`
}
