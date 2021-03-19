package models

type Product struct {
	ID      int64
	StoreID int64
	Title   string

	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

type FullProduct struct {
}
