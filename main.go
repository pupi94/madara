package main

import (
	"context"
	"os"

	"github.com/pupi94/madara/cmd"
	"github.com/pupi94/madara/config"
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
				Usage: "migrate database. db:migrate up OR db:migrate down. Default up",
				Action: func(c *cli.Context) error {
					var direction = c.Args().First()
					config.InitDB()
					return cmd.DbMigrate(ctx, direction, c.Int("step"))
				},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "step",
						Usage: "db migrate up or down step",
						Value: 1,
					},
				},
			},
			{
				Name:  "db:generate_migration",
				Usage: "generate migration",
				Action: func(c *cli.Context) error {
					var name = c.Args().First()
					config.InitDB()
					return cmd.GenerateMigration(ctx, name)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.WithError(err).Fatal("App start fail....")
	}
}
