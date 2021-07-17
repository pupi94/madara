package model

import (
	_ "github.com/Masterminds/squirrel"
	_ "github.com/jmoiron/sqlx"
	"github.com/pupi94/madara/component/xtype"
)

type Collect struct {
	ID           xtype.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	CollectionID xtype.Uuid     `db:"collection_id" json:"collection_id" faker:"uuid_hyphenated"`
	ProductID    xtype.Uuid     `db:"product_id" json:"product_id" faker:"uuid_hyphenated"`
	StoreID      xtype.Uuid     `db:"store_id" json:"store_id" faker:"uuid_hyphenated"`
	Position     *int64         `db:"position" json:"position" faker:"int64"`
	CreatedAt    *xtype.UtcTime `db:"created_at" json:"created_at" faker:"-"`
	UpdatedAt    *xtype.UtcTime `db:"updated_at" json:"updated_at" faker:"-"`
}

func (ct Collect) TableName() string {
	return "collects"
}
