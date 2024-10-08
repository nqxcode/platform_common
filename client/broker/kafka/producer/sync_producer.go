package producer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"

	"github.com/nqxcode/platform_common/client/broker/kafka"
)

type syncProducer struct {
	producer sarama.SyncProducer
}

// ProduceResult produce result
type ProduceResult struct {
	Partition int32
	Offset    int64
}

// NewSyncProducer new sync producer
func NewSyncProducer(cfg kafka.ProducerConfig) (*syncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = cfg.RequiredAcks()
	config.Producer.Retry.Max = cfg.RetryMax()
	config.Producer.Return.Successes = cfg.ReturnSuccesses()

	producer, err := sarama.NewSyncProducer(cfg.Brokers(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to start producer: %w", err)
	}

	return &syncProducer{
		producer: producer,
	}, nil
}

// Produce produces message
func (p *syncProducer) Produce(_ context.Context, topicName string, value any) (int32, int64, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to marshal message: %v", err)
	}

	produceMessage := &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.StringEncoder(data),
	}

	partition, offset, err := p.producer.SendMessage(produceMessage)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to send message in Kafka: %w", err)
	}

	return partition, offset, nil
}

// Close close producer
func (p *syncProducer) Close() error {
	err := p.producer.Close()
	if err != nil {
		return fmt.Errorf("failed to close producer: %w", err)
	}

	return nil
}
