package interceptor

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"path"
	"time"
)

func AccessLogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if _, ok := info.Server.(grpc_health_v1.HealthServer); ok {
			return handler(ctx, req)
		}
		startTime := time.Now()
		resp, err = handler(ctx, req)

		method := path.Base(info.FullMethod)
		grpcCode := status.Code(err)
		logrus.WithField("method", method).WithField("Code", grpcCode).Info("Time: ", time.Now().Sub(startTime))
		return resp, err
	}
}