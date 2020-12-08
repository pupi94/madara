package consumers

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func ConsumeImage(ctx context.Context, m *sarama.ConsumerMessage) error {
	logrus.Info("Value = ", string(m.Value), ";  Partition = ", m.Partition, "   Key = ", m.Key)
	return nil
}
