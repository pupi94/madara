package models

type Store struct {
	ID        uint64 `gorm:"primaryKey"`
	CreatedAt int64  `gorm:"autoCreateTime`
	UpdatedAt int64  `gorm:"autoUpdateTime`
}
