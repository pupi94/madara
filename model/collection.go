package model

import (
	"github.com/pupi94/madara/component/xtype"
)

type Collection struct {
	ID        xtype.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	StoreID   xtype.Uuid     `db:"store_id" json:"store_id" faker:"uuid_hyphenated"`
	Title     *string        `db:"title" json:"title" faker:"word"`
	CreatedAt *xtype.UtcTime `db:"created_at" json:"created_at" faker:"-"`
	UpdatedAt *xtype.UtcTime `db:"updated_at" json:"updated_at" faker:"-"`
}

func (Collection) TableName() string {
	return "collections"
}
