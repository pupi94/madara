package model

import "github.com/pupi94/madara/component/xtype"

type Store struct {
	ID        xtype.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	CreatedAt *xtype.UtcTime `db:"created_at" json:"created_at" faker:"-"`
	UpdatedAt *xtype.UtcTime `db:"updated_at" json:"updated_at" faker:"-"`
}

func (Store) TableName() string {
	return "stores"
}
