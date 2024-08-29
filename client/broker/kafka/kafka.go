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
