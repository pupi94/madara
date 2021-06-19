package models

import (
	"github.com/pupi94/madara/components/xtypes"
)

type Product struct {
	ID                xtypes.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	StoreID           xtypes.Uuid     `db:"store_id" json:"store_id" faker:"uuid_hyphenated"`
	Title             *string         `db:"title" json:"title" faker:"word"`
	Description       *string         `db:"description" json:"description" faker:"sentence"`
	PublishedAt       *xtypes.UtcTime `db:"published_at" json:"published_at" faker:"-"`
	InventoryQuantity *int64          `db:"inventory_quantity" json:"inventory_quantity" faker:"-"`
	CreatedAt         *xtypes.UtcTime `db:"created_at" json:"created_at" faker:"-"`
	UpdatedAt         *xtypes.UtcTime `db:"updated_at" json:"updated_at" faker:"-"`
	Published         *bool           `db:"published" json:"published" faker:"false"`
}

func (p Product) TableName() string {
	return "products"
}

type FullProduct struct {
	Product
	Images   []*Image
	Variants []*Variant
}
