package producer

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/nqxcode/platform_common/client/broker/kafka"
)

type syncProducer struct {
	config   kafka.ProducerConfig
	producer sarama.SyncProducer
}

type Message struct {
	Topic string
	Value any
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
func (p *syncProducer) Produce(message *Message) (*ProduceResult, error) {
	if message == nil {
		return nil, fmt.Errorf("message is nil")
	}

	data, err := json.Marshal(message.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %v", err)
	}

	produceMessage := &sarama.ProducerMessage{
		Topic: message.Topic,
		Value: sarama.StringEncoder(data),
	}

	partition, offset, err := p.producer.SendMessage(produceMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to send message in Kafka: %w", err)
	}

	return &ProduceResult{Partition: partition, Offset: offset}, nil
}

// Close close producer
func (p *syncProducer) Close() error {
	err := p.producer.Close()
	if err != nil {
		return fmt.Errorf("failed to close producer: %w", err)
	}

	return nil
}
