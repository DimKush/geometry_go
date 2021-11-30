package repository

import (
	"gorm.io/gorm"

	"github.com/DimKush/geometry_go/internal/entity/warehouse"
	"github.com/DimKush/geometry_go/internal/entity/unit"
)

type Warehouse interface {
	GetWarehouseById(id int) (*warehouse.ItemWarehouse, error)
}

type Unit interface {
	SetUnit(unit unit.Unit) error
}

type Repository struct {
	Warehouse
	Unit
}

func InitRepository(db *gorm.DB) *Repository {
	return &Repository{
		Warehouse: InitWarehouseRep(db),
		Unit: InitUnitRep(db),
	}
}