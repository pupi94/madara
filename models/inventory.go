package models

import (
	"context"
	"github.com/pupi94/madara/components/xtypes"
	"gorm.io/gorm"
)

type Inventory struct {
	ID        xtypes.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	ProductID xtypes.Uuid     `db:"source_id" json:"source_id" faker:"uuid_hyphenated"`
	Value     *int64          `db:"value" json:"value" faker:"int64"`
	CreatedAt *xtypes.UtcTime `db:"created_at" json:"created_at"`
	UpdatedAt *xtypes.UtcTime `db:"updated_at" json:"updated_at"`
}

func (Inventory) TableName() string {
	return "inventories"
}

func GetProductInventory(ctx context.Context, db *gorm.DB, productID xtypes.Uuid) (int64, error) {
	return 0, nil
}
