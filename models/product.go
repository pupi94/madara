package models

type Product struct {
	ID          uint64 `gorm:"primaryKey"`
	StoreID     uint64 `gorm:"not null"`
	Title       string
	Description string
	Published   bool `gorm:"default:false"`
	PublishedAt *int64
	CreatedAt   int64 `gorm:"autoCreateTime`
	UpdatedAt   int64 `gorm:"autoUpdateTime`
}

type FullProduct struct {
}
