package kafka

import (
	"context"
	"github.com/IBM/sarama"
)

type ConsumerGroup struct {
	sarama.ConsumerGroup

	groupID string
	topics  []string
}

type ConsumerGroupConfig struct {
	KafkaVersion   sarama.KafkaVersion
	OffsetsInitial int64
	IsReturnErr    bool
}

func NewConsumerGroup(consumerConfig *ConsumerGroupConfig, topics, addr []string, groupID string) *ConsumerGroup {
	config := sarama.NewConfig()
	config.Version = consumerConfig.KafkaVersion
	config.Consumer.Offsets.Initial = consumerConfig.OffsetsInitial
	config.Consumer.Return.Errors = consumerConfig.IsReturnErr
	client, err := sarama.NewClient(addr, config)
	if err != nil {
		panic(err)
	}

	consumerGroup, err := sarama.NewConsumerGroupFromClient(groupID, client)
	if err != nil {
		panic(err)
	}

	g := &ConsumerGroup{
		consumerGroup,
		groupID,
		topics,
	}
	return g
}

func (cg *ConsumerGroup) RegisterHandleAndConsumer(handler sarama.ConsumerGroupHandler) {
	ctx := context.Background()
	for {
		err := cg.ConsumerGroup.Consume(ctx, cg.topics, handler)
		if err != nil {
			panic(err.Error())
		}
	}
}
