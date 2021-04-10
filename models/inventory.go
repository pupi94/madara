package models

import "github.com/pupi94/madara/components/xtypes"

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
