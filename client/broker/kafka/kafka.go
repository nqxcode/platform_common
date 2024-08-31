package kafka

import (
	"context"

	"github.com/nqxcode/platform_common/client/broker/kafka/consumer"
	"github.com/nqxcode/platform_common/client/broker/kafka/producer"
)

// Consumer kafka consumer interface
type Consumer interface {
	Consume(ctx context.Context, topicName string, handler consumer.Handler) (err error)
	Close() error
}

// SyncProducer sync producer interface
type SyncProducer interface {
	Produce(message *producer.Message) (*producer.ProduceResult, error)
	Close() error
}
