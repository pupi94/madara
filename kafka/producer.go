package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type Producer interface {
	Produce(topic string, message []byte, key string) error
}

type SyncProducer struct {
	producer sarama.SyncProducer
}

type AsyncProducer struct {
	producer  sarama.AsyncProducer
	failTimes int32
}

func NewSyncProducer(hosts []string) (Producer, error) {
	var sp = new(SyncProducer)

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	// 分区选择机制
	cfg.Producer.Partitioner = sarama.NewReferenceHashPartitioner

	var err error
	sp.producer, err = sarama.NewSyncProducer(hosts, cfg)
	if err != nil {
		logrus.WithError(err).Panic("create kafka async producer failed")
	}
	return sp, err
}

func (sp *SyncProducer) Produce(topic string, message []byte, key string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
		Key:   sarama.StringEncoder(key),
	}

	_, _, err := sp.producer.SendMessage(msg)
	return err
}

func NewAsyncProducer(ctx context.Context, hosts []string) (Producer, error) {
	var ap = new(AsyncProducer)

	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	// 分区选择机制
	cfg.Producer.Partitioner = sarama.NewReferenceHashPartitioner

	var err error
	ap.producer, err = sarama.NewAsyncProducer(hosts, cfg)
	if err != nil {
		logrus.WithError(err).Panic("create kafka async producer failed")
	}
	go ap.Run(ctx)
	return ap, err
}

func (ap *AsyncProducer) Produce(topic string, message []byte, key string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
		Key:   sarama.StringEncoder(key),
	}
	ap.producer.Input() <- msg
	return nil
}

func (ap *AsyncProducer) Run(ctx context.Context) {
	defer ap.producer.AsyncClose()
	for {
		select {
		case <-ap.producer.Successes():
		case fail := <-ap.producer.Errors():
			ap.failTimes += 1

			val, _ := fail.Msg.Value.Encode()
			logrus.WithFields(logrus.Fields{
				"topic":      fail.Msg.Topic,
				"partitions": fail.Msg.Partition,
				"value":      string(val),
				"failTimes":  ap.failTimes,
			}).WithError(fail.Err).Warn("send kafka failed")
		case <-ctx.Done():
			return
		}
	}
}
