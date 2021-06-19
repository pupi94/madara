package models

import "github.com/pupi94/madara/components/xtypes"

type Store struct {
	ID        xtypes.Uuid     `db:"id" json:"id" faker:"uuid_hyphenated"`
	CreatedAt *xtypes.UtcTime `db:"created_at" json:"created_at" faker:"-"`
	UpdatedAt *xtypes.UtcTime `db:"updated_at" json:"updated_at" faker:"-"`
}

func (Store) TableName() string {
	return "stores"
}
