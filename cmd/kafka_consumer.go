package cmd

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/pupi94/madara/config"
	"github.com/pupi94/madara/kafka/consumers"
	"github.com/pupi94/madara/tools/kafka"
)

var GroupId = "madara"

func ConsumeImage(ctx context.Context) error {
	brokers := config.Env.KafkaHost
	topics := []string{config.Env.ImageTopicName}

	return kafka.Consume(topics, brokers, GroupId, func(message *sarama.ConsumerMessage) error {
		return consumers.ConsumeImage(ctx, message)
	})
}
