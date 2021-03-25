package models

type Inventory struct {
	ID        uint64
	StoreID   uint64
	ProductID uint64
	Value     uint64
	CreatedAt int64 `gorm:"autoCreateTime`
	UpdatedAt int64 `gorm:"autoUpdateTime`
}
