package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
	//producer  sarama.AsyncProducer
	failTimes int32
}

func NewKafkaProducer(ctx context.Context, hosts []string) (*KafkaProducer, error) {
	var kp = new(KafkaProducer)

	cfg := sarama.NewConfig()

	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.Partitioner = sarama.NewHashPartitioner

	var err error
	//kp.producer, err = sarama.NewAsyncProducer(hosts, cfg)
	kp.producer, err = sarama.NewSyncProducer(hosts, cfg)
	if err != nil {
		logrus.WithError(err).Panic("create kafka async producer failed")
	}
	//go kp.Run(ctx)
	return kp, err
}

func (kp *KafkaProducer) Produce(topic string, message []byte, partition int32) error {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.ByteEncoder(message),
		Partition: partition,
	}

	partition, offset, err := kp.producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	logrus.Info("Producer partition = ", partition, "  offset = ", offset)
	//kp.producer.Input() <- msg
	return err
}

/*
func (kp *KafkaProducer) Run(ctx context.Context) {
	defer kp.producer.AsyncClose()
	for {
		select {
		case <-kp.producer.Successes():
		case fail := <-kp.producer.Errors():
			kp.failTimes += 1

			val, _ := fail.Msg.Value.Encode()
			logrus.WithFields(logrus.Fields{
				"topic":      fail.Msg.Topic,
				"partitions": fail.Msg.Partition,
				"value":      string(val),
				"failTimes":  kp.failTimes,
			}).WithError(fail.Err).Warn("send kafka failed")
		case <-ctx.Done():
			return
		}
	}
}*/
