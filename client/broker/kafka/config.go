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

// ProducerConfig kafka producer configuration
type ProducerConfig interface {
	Brokers() []string
	RequiredAcks() sarama.RequiredAcks
	RetryMax() int
	ReturnSuccesses() bool
}
