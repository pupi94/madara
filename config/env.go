package config

import (
	"github.com/caarlos0/env/v6"
)

type Environment struct {
	GrpcPort     int      `env:"GrpcPort" envDefault:"3000"`
	DBHostname   string   `env:"DB_HOSTNAME" envDefault:"127.0.0.1"`
	DBPort       int      `env:"DB_PORT" envDefault:"3306"`
	DBUsername   string   `env:"DB_USERNAME" envDefault:"root"`
	DBPassword   string   `env:"DB_PASSWORD" envDefault:"123456"`
	DBDatabase   string   `env:"DB_DATABASE" envDefault:"product_development"`
}

var Env *Environment

func init() {
	Env = &Environment{}
	if err := env.Parse(Env); err != nil {
		panic(err)
	}
}
