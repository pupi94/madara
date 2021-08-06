package product

import (
	. "github.com/smartystreets/goconvey/convey"
)

func (suite *ProductSuite) TestCreateProduct() {
	form := CreateForm{
		Title:             "",
		Published:         true,
		Description:       "dsdsdsdsd",
		InventoryQuantity: 12,
		Variants: []*CreateFormVariant{
			{Price: 12, CompareAtPrice: 21, Barcode: "dsdsd", InventoryQuantity: 12},
		},
		Images: []*CreateFormImage{
			{Src: "", Alt: "dsdsd", Width: 32, Height: 32},
		},
	}

	Convey("create 1 product", suite.T(), func() {
		err := form.Validate()
		if err != nil {
			panic(err)
		}
	})
}
