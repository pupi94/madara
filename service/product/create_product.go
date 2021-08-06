package product

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/pupi94/madara/model"
)

type CreateFormVariant struct {
	Price             float64 `valid:"Required"`
	CompareAtPrice    float64 `valid:"Required"`
	Barcode           string
	InventoryQuantity int64
}

type CreateFormImage struct {
	Src    string `valid:"Required"`
	Alt    string
	Width  int32
	Height int32
}

type CreateForm struct {
	Title             string `valid:"Required;MaxSize(255)"`
	Description       string
	Published         bool
	InventoryQuantity int64
	Variants          []*CreateFormVariant
	Images            []*CreateFormImage
}

func (f *CreateForm) Valid(v *validation.Validation) {
	if len(f.Variants) > 200 {
		v.SetError("Variants", "variants_size_over_limit")
		return
	}
	if len(f.Images) > 100 {
		v.SetError("Images", "images_size_over_limit")
		return
	}
}

func (f *CreateForm) Validate() error {
	valid := validation.Validation{}
	ok, err := valid.Valid(f)
	if err != nil {
		return err
	}
	if !ok {
		for _, e := range valid.Errors {
			fmt.Printf("%#v", e)
		}
	}
	return nil
}

func CreateProduct(ctx context.Context, form *CreateForm) (*model.FullProduct, error) {
	if err := form.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}
