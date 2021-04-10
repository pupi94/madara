package models

import (
	"context"
	"database/sql"
	"github.com/pupi94/madara/components/xtypes"
	"time"

	sq "github.com/Masterminds/squirrel"
	"gitlab.shoplazza.site/shoplaza/samoyed/goat/db"

	"github.com/jmoiron/sqlx"
)

type ProductCollection struct {
	ID           xtypes.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	CollectionID xtypes.Uuid     `db:"collection_id" json:"collection_id" faker:"uuid_hyphenated"`
	ProductID    xtypes.Uuid     `db:"product_id" json:"product_id" faker:"uuid_hyphenated"`
	CreatedAt    *xtypes.UtcTime `db:"created_at" json:"created_at" faker:"-"`
	UpdatedAt    *xtypes.UtcTime `db:"updated_at" json:"updated_at" faker:"-"`
	Position     *int64          `db:"position" json:"position" faker:"int64"`
	StoreID      xtypes.Uuid     `db:"store_id" json:"store_id" faker:"uuid_hyphenated"`
}

func (ProductCollection) TableName() string {
	return "product_collections"
}
