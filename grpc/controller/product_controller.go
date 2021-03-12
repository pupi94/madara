package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pupi94/madara/grpc/pb"
)

type ProductController struct {
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (controller *ProductController) CreateProduct(ctx context.Context, request *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	return nil, nil
}

func (controller *ProductController) DeleteProduct(ctx context.Context, request *pb.DeleteProductRequest) (*empty.Empty, error) {
	return nil, nil
}

func (controller *ProductController) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	return nil, nil
}

func (controller *ProductController) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	return nil, nil
}
func (controller *ProductController) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductResponse, error) {
	return nil, nil
}
