package cmd

import (
	"context"
	"github.com/pupi94/madara/cmd/grpc"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := &cli.App{
		Name: "grpc",
		Usage: "Start grpc service",
		Action: func(c *cli.Context) error {
			_ = grpc.StartGrpcServer(ctx)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.WithError(err).Fatal("App start fail....")
	}
}
