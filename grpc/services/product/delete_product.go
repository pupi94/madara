package product

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pupi94/madara/grpc/pb"
)

func (ser *Service) DeleteProduct(ctx context.Context, request *pb.DeleteProductRequest) (*empty.Empty, error) {
	return nil, nil
}
