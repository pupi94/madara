package model

import (
	"context"
	"github.com/pupi94/madara/component/xtype"
	"gorm.io/gorm"
)

type Inventory struct {
	ID         xtype.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	StoreID    xtype.Uuid     `db:"store_id" json:"store_id" faker:"uuid_hyphenated"`
	SourceType string         `db:"source_type" json:"source_type" faker:"word"`
	SourceID   xtype.Uuid     `db:"source_id" json:"source_id" faker:"uuid_hyphenated"`
	Value      *int64         `db:"value" json:"value" faker:"int64"`
	CreatedAt  *xtype.UtcTime `db:"created_at" json:"created_at"`
	UpdatedAt  *xtype.UtcTime `db:"updated_at" json:"updated_at"`
}

func (Inventory) TableName() string {
	return "inventories"
}

func GetInventory(ctx context.Context, db *gorm.DB, sourceId xtype.Uuid) (int64, error) {
	return 0, nil
}

func SelectInventories(ctx context.Context, db *gorm.DB, sourceIds []xtype.Uuid) ([]*Inventory, error) {
	return nil, nil
}
