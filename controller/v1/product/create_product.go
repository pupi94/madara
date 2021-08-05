package product

import (
	"context"
	v1 "github.com/pupi94/madara/grpc/pb/v1"
	"github.com/pupi94/madara/model"
	service "github.com/pupi94/madara/service/product"
)

func (controller *ProductController) CreateProduct(ctx context.Context, req *v1.CreateProductRequest) (*v1.ProductResponse, error) {
	form := buildCreateProductForm(req)
	product, err := service.CreateProduct(ctx, form)
	if err != nil {
		return nil, err
	}
	return buildProductResponse(product), nil
}

func buildCreateProductForm(request *v1.CreateProductRequest) *service.CreateForm {
	var variants []*service.CreateFormVariant
	for _, v := range request.GetVariants() {
		variants = append(variants, &service.CreateFormVariant{
			Price:             v.GetPrice(),
			CompareAtPrice:    v.GetCompareAtPrice(),
			Barcode:           v.GetBarcode(),
			InventoryQuantity: v.GetInventoryQuantity(),
		})
	}
	var images []*service.CreateFormImage
	for _, img := range request.GetImages() {
		images = append(images, &service.CreateFormImage{
			Src:    img.GetSrc(),
			Alt:    img.GetAlt(),
			Width:  img.GetWidth(),
			Height: img.GetHeight(),
		})
	}
	form := &service.CreateForm{
		Title:       request.GetTitle(),
		Description: request.GetDescription(),
		Published:   request.GetPublished(),
		Variants:    variants,
		Images:      images,
	}
	return form
}

func buildProductResponse(product *model.FullProduct) *v1.ProductResponse {
	return nil
}
