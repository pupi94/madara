package server

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pupi94/madara/controller/v1"
	"github.com/pupi94/madara/grpc/interceptor"
	pb_v1 "github.com/pupi94/madara/grpc/pb/v1"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	server *grpc.Server
}

func NewGrpcServer() *GrpcServer {
	opts := []grpc.ServerOption{grpcInterceptors()}

	sev := grpc.NewServer(opts...)
	return &GrpcServer{server: sev}
}

func (rp *GrpcServer) Start(port int) error {
	rp.registerService()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return rp.server.Serve(lis)
}

func (rp *GrpcServer) registerService() {
	pb_v1.RegisterProductControllerServer(rp.server, v1.NewProductController())
}

func (rp *GrpcServer) Shutdown(ctx context.Context) (err error) {
	ch := make(chan struct{})
	go func() {
		rp.server.GracefulStop()
		close(ch)
	}()
	select {
	case <-ctx.Done():
		rp.server.Stop()
		err = ctx.Err()
	case <-ch:
	}
	return
}

func grpcInterceptors() grpc.ServerOption {
	return grpc_middleware.WithUnaryServerChain(
		interceptor.AccessLogInterceptor(),
	)
}
