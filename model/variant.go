package model

import (
	"github.com/pupi94/madara/component/xtype"
	"github.com/shopspring/decimal"
)

type Variant struct {
	ID                xtype.Uuid           `db:"id" json:"id" faker:"uuid_hyphenated"`
	StoreID           xtype.Uuid           `db:"store_id" json:"store_id" faker:"uuid_hyphenated"`
	ProductID         xtype.Uuid           `db:"product_id" json:"product_id" faker:"uuid_hyphenated"`
	Position          *int64               `db:"position" json:"position" faker:"int64"`
	Title             *string              `db:"title" json:"title" faker:"word"`
	CompareAtPrice    *decimal.NullDecimal `db:"compare_at_price" json:"compare_at_price" faker:"decimal"`
	Price             *decimal.NullDecimal `db:"price" json:"price" faker:"decimal"`
	Barcode           *string              `db:"barcode" json:"barcode" faker:"word"`
	InventoryQuantity *int64               `db:"inventory_quantity" json:"inventory_quantity" faker:"-"`
	Weight            *decimal.NullDecimal `db:"weight" json:"weight" faker:"decimal"`
	WeightUnit        *string              `db:"weight_unit" json:"weight_unit" faker:"weight_unit"`
	CreatedAt         *xtype.UtcTime       `db:"created_at" json:"created_at" faker:"-"`
	UpdatedAt         *xtype.UtcTime       `db:"updated_at" json:"updated_at" faker:"-"`
	ImageID           xtype.Uuid           `db:"image_id" json:"image_id" faker:"uuid_hyphenated"`
	Note              *string              `db:"note" json:"note" faker:"word"`
}

func (v Variant) TableName() string {
	return "variants"
}
