package model

import "github.com/pupi94/madara/component/xtype"

type Image struct {
	ID         xtype.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	StoreID    xtype.Uuid     `db:"store_id" json:"store_id" faker:"uuid_hyphenated"`
	ProductID  xtype.Uuid     `db:"product_id" json:"store_id" faker:"uuid_hyphenated"`
	Position   int64          `db:"position" json:"position""`
	Properties *xtype.JSON    `db:"properties" json:"properties" faker:"image"`
	CreatedAt  *xtype.UtcTime `db:"created_at" json:"created_at"`
	UpdatedAt  *xtype.UtcTime `db:"updated_at" json:"updated_at"`
}

func (img Image) TableName() string {
	return "images"
}
