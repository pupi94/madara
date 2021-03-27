package es

type Product struct {
	ID                uint64 `json:"id"`
	StoreID           uint64 `json:"store_id"`
	Title             string `json:"title"`
	Published         bool   `json:"published"`
	CreatedAt         int64  `json:"created_at"`
	UpdatedAt         int64  `json:"updated_at"`
	InventoryQuantity int64  `json:"inventory_quantity"`
}

func IndexProduct(product *Product) error {
	return nil
}

func BatchIndexProduct(products []*Product) error {
	return nil
}

func DeleteProduct(productID uint) error {
	return nil
}
