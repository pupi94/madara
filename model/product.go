package model

import (
	"context"
	"github.com/pupi94/madara/component/xtype"
	"gorm.io/gorm"
)

type Product struct {
	ID                xtype.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	StoreID           xtype.Uuid     `db:"store_id" json:"store_id" faker:"uuid_hyphenated"`
	Title             *string        `db:"title" json:"title" faker:"word"`
	Description       *string        `db:"description" json:"description" faker:"sentence"`
	Published         *bool          `db:"published" json:"published" faker:"false"`
	PublishedAt       *xtype.UtcTime `db:"published_at" json:"published_at" faker:"-"`
	InventoryQuantity *int64         `db:"inventory_quantity" json:"inventory_quantity" faker:"-"`
	CreatedAt         *xtype.UtcTime `db:"created_at" json:"created_at" faker:"-"`
	UpdatedAt         *xtype.UtcTime `db:"updated_at" json:"updated_at" faker:"-"`
}

func (p Product) TableName() string {
	return "products"
}

type FullProduct struct {
	Product
	Images   []*Image   `json:"images"`
	Variants []*Variant `json:"variants"`
}

func GetProductDescription(ctx context.Context, db *gorm.DB, storeId, id xtype.Uuid) (string, error) {
	return "", nil
}
