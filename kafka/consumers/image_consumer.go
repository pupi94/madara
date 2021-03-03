package consumers

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type Image struct {
	StoreID string `json:"store_id"`
	Host    string `json:"host"`
	Src     string `json:"src"`
}

func ConsumeImage(ctx context.Context, m *sarama.ConsumerMessage) error {
	var img Image
	if err := json.Unmarshal(m.Value, &img); err != nil {
		logrus.Panic("Unmarshal Failed: ", err)
	}

	logrus.Info("Partition: ", m.Partition, "   Store-id: ", img.StoreID, "    Host: ", img.Host)
	return nil
}
