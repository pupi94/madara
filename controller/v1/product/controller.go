package product

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pupi94/madara/grpc/pb/v1"
)

type ProductController struct {
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (controller *ProductController) DeleteProduct(ctx context.Context, request *v1.DeleteProductRequest) (*empty.Empty, error) {
	return nil, nil
}

func (controller *ProductController) GetProduct(ctx context.Context, req *v1.GetProductRequest) (*v1.ProductResponse, error) {
	return nil, nil
}
func (controller *ProductController) ListProduct(ctx context.Context, req *v1.ListProductRequest) (*v1.ListProductResponse, error) {
	return nil, nil
}
