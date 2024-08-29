package kafka

import (
	"github.com/IBM/sarama"
)

type ConsumerConfig interface {
	Brokers() []string
	GroupID() string
	Config() *sarama.Config
}
