package models

import (
	_ "github.com/Masterminds/squirrel"
	_ "github.com/jmoiron/sqlx"
	"github.com/pupi94/madara/components/xtypes"
)

type Collect struct {
	ID           xtypes.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	CollectionID xtypes.Uuid     `db:"collection_id" json:"collection_id" faker:"uuid_hyphenated"`
	ProductID    xtypes.Uuid     `db:"product_id" json:"product_id" faker:"uuid_hyphenated"`
	StoreID      xtypes.Uuid     `db:"store_id" json:"store_id" faker:"uuid_hyphenated"`
	Position     *int64          `db:"position" json:"position" faker:"int64"`
	CreatedAt    *xtypes.UtcTime `db:"created_at" json:"created_at" faker:"-"`
	UpdatedAt    *xtypes.UtcTime `db:"updated_at" json:"updated_at" faker:"-"`
}

func (ct Collect) TableName() string {
	return "collects"
}
