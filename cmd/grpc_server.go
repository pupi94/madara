package cmd

import (
	"context"
	"github.com/pupi94/madara/config"
	"github.com/pupi94/madara/grpc/server"
	"github.com/sirupsen/logrus"
)

func StartGrpcServer(ctx context.Context) error {
	sev := server.NewGrpcServer()
	err := sev.Start(config.Env.GrpcPort)
	if err != nil {
		logrus.WithError(err).Error("grpc server shutdown failed")
		return err
	}

	select {
	case <-ctx.Done():
	}

	err = sev.Shutdown(context.TODO())
	if err != nil {
		logrus.WithError(err).Error("grpc server shutdown failed")
		return err
	}
	logrus.Info("grpc server exited")
	return nil
}
