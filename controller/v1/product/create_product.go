package product

import (
	"context"
	v1 "github.com/pupi94/madara/grpc/pb/v1"
)

func (controller *ProductController) CreateProduct(ctx context.Context, request *v1.CreateProductRequest) (*v1.ProductResponse, error) {
	return &v1.ProductResponse{Id: "xxxxxx"}, nil
}
