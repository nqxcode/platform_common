package kafka

import (
	"github.com/IBM/sarama"
)

// ConsumerConfig kafka consumer configuration
type ConsumerConfig interface {
	Brokers() []string
	GroupID() string
	Config() *sarama.Config
}
