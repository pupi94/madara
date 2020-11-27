package grpc

import (
	"context"
	"github.com/pupi94/madara/config"
	"github.com/pupi94/madara/grpc/pb"
	"github.com/pupi94/madara/services/product"
	"log"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func StartGrpcServer(ctx context.Context) error {
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			//RecoveryInterceptor,
			//LoggingInterceptor,
		),
	}

	server := grpc.NewServer(opts...)
	registerService(server)

	lis, err := net.Listen("tcp", ":"+config.Env.GrpcPort)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	return server.Serve(lis)
}

func registerService(server *grpc.Server)  {
	pb.RegisterProductServiceServer(server, product.NewService())
}
