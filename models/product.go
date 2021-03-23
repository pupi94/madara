package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	StoreID     uuid.UUID `gorm:"type:uuid;not null"`
	Title       string
	Description string
	Published   bool `gorm:"default:false"`
	PublishedAt int64
	CreatedAt   int64 `gorm:"autoCreateTime`
	UpdatedAt   int64 `gorm:"autoUpdateTime`
}

type FullProduct struct {
}
