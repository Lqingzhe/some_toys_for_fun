package commonconfig

import (
	"github.com/IBM/sarama"
)

func GetKafkaBasicConfig() sarama.Config {
	config := sarama.Config{}
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	return config
}
