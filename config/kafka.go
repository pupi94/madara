package config

import (
	"context"
	"github.com/pupi94/madara/kafka"
)

var AsyncProducer *kafka.KafkaProducer

func InitKafkaProducer(ctx context.Context) {
	var err error
	AsyncProducer, err = kafka.NewKafkaProducer(ctx, Env.KafkaHost)
	if err != nil {
		panic(err)
	}
}
