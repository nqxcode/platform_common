package kafka

import (
	"context"

	"github.com/nqxcode/platform_common/client/broker/kafka/consumer"
)

// Consumer kafka consumer interface
type Consumer interface {
	Consume(ctx context.Context, topicName string, handler consumer.Handler) (err error)
	Close() error
}

// SyncProducer sync producer interface
type SyncProducer interface {
	Produce(ctx context.Context, topicName string, data any) (partition int32, offset int64, err error)
	Close() error
}
