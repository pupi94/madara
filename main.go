package main

import (
	"context"
	"github.com/pupi94/madara/cmd"
	"github.com/pupi94/madara/config"
	"github.com/pupi94/madara/db"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "grpc",
				Usage: "Start grpc service",
				Action: func(c *cli.Context) error {
					config.InitDB()
					config.InitKafkaProducer(ctx)
					return cmd.StartGrpcServer(ctx)
				},
			},
			{
				Name:  "consume_image",
				Usage: "consume image",
				Action: func(c *cli.Context) error {
					return cmd.ConsumeImage(ctx)
				},
			},
			{
				Name:  "db:migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					config.InitDB()
					db.Migrate()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.WithError(err).Fatal("App start fail....")
	}
}
