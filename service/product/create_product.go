package product

import (
	"context"
	"github.com/pupi94/madara/model"
)

type CreateFormVariant struct {
	Price             float64
	CompareAtPrice    float64
	Barcode           string
	InventoryQuantity int64
}

type CreateFormImage struct {
	Src    string
	Alt    string
	Width  int32
	Height int32
}

type CreateForm struct {
	Title             string
	Description       string
	Published         bool
	InventoryQuantity int64
	Variants          []*CreateFormVariant
	Images            []*CreateFormImage
}

func CreateProduct(ctx context.Context, form *CreateForm) (*model.FullProduct, error) {
	return nil, nil
}
