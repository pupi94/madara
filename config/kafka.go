package config

import (
	"context"
	"github.com/pupi94/madara/component/kafka"
)

var AsyncProducer kafka.Producer

func InitKafkaProducer(ctx context.Context) {
	var err error
	AsyncProducer, err = kafka.NewSyncProducer(Env.KafkaHost)
	if err != nil {
		panic(err)
	}
}
