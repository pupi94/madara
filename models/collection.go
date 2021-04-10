package models

import (
	"github.com/pupi94/madara/components/xtypes"
)

type Collection struct {
	ID        xtypes.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	StoreID   xtypes.Uuid     `db:"store_id" json:"store_id" faker:"uuid_hyphenated"`
	Title     *string         `db:"title" json:"title" faker:"word"`
	CreatedAt *xtypes.UtcTime `db:"created_at" json:"created_at" faker:"-"`
	UpdatedAt *xtypes.UtcTime `db:"updated_at" json:"updated_at" faker:"-"`
}

func (Collection) TableName() string {
	return "collections"
}
