package models

import "github.com/pupi94/madara/components/xtypes"

type Image struct {
	ID         xtypes.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	StoreID    xtypes.Uuid     `db:"store_id" json:"store_id" faker:"uuid_hyphenated"`
	Properties *xtypes.JSON    `db:"properties" json:"properties" faker:"image"`
	CreatedAt  *xtypes.UtcTime `db:"created_at" json:"created_at"`
	UpdatedAt  *xtypes.UtcTime `db:"updated_at" json:"updated_at"`
}

func (img Image) TableName() string {
	return "images"
}
