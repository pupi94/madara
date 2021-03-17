package config

import (
	"github.com/caarlos0/env/v6"
)

type Environment struct {
	GrpcPort int `env:"GrpcPort" envDefault:"3000"`

	DBHostname string `env:"DB_HOSTNAME" envDefault:"127.0.0.1"`
	DBPort     int    `env:"DB_PORT" envDefault:"3306"`
	DBUsername string `env:"DB_USERNAME" envDefault:"root"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"123456"`
	DBDatabase string `env:"DB_DATABASE" envDefault:"madara_development"`
	DBShowSQL  bool   `env:"DB_SHOW_SQL" envDefault:"true"`

	RdsHost      string `env:"RDS_HOST" envDefault:"127.0.0.1"`
	RdsPort      int    `env:"RDS_PORT" envDefault:"6379"`
	RdsNamespace string `env:"RDS_NAMESPACE" envDefault:""`

	KafkaHost      []string `env:"KAFKA_SERVER"  envDefault:"127.0.0.1:9092"`
	ImageTopicName string   `env:"IMAGE_TOPIC_NAME"  envDefault:"image"`

	LogLevel string `env:"LOG_LEVEL" envDefault:"info"` // panic fatal error warn warning info debug trace
}

var Env *Environment

func init() {
	Env = &Environment{}
	if err := env.Parse(Env); err != nil {
		panic(err)
	}
}
