package kafka

import (
	"github.com/IBM/sarama"
	"sync"
)

type Consumer struct {
	addr          []string
	Topic         string
	PartitionList []int32
	Consumer      sarama.Consumer
	WG            sync.WaitGroup
}

func NewKafkaConsumer(addr []string, topic string) *Consumer {
	consumer, err := sarama.NewConsumer(addr, nil)
	if err != nil {
		panic(err)
		return nil
	}

	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
		return nil
	}

	p := &Consumer{
		addr:          addr,
		Topic:         topic,
		PartitionList: partitionList,
		Consumer:      consumer,
	}
	return p
}
