package product

import (
	"context"
	"github.com/pupi94/madara/grpc/pb"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (ser *Service) UpdateProduct(ctx context.Context, request *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	panic("implement me")
}

func (ser *Service) GetProduct(ctx context.Context, request *pb.GetProductRequest) (*pb.ProductResponse, error) {
	panic("implement me")
}

func (ser *Service) SearchProducts(ctx context.Context, request *pb.SearchProductsRequest) (*pb.SearchProductsResponse, error) {
	panic("implement me")
}