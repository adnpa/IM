package kafka

import (
	"github.com/IBM/sarama"
	"github.com/gogo/protobuf/proto"
)

type Producer struct {
	topic    string
	addr     []string
	config   *sarama.Config
	producer sarama.SyncProducer
}

func NewProducer(addr []string, topic string) *Producer {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Partitioner = sarama.NewHashPartitioner

	producer, err := sarama.NewSyncProducer(addr, cfg)
	if err != nil {
		panic(err)
		return nil
	}

	p := &Producer{
		topic:    topic,
		addr:     addr,
		config:   cfg,
		producer: producer,
	}
	return p
}

func (p *Producer) SendMsg(m proto.Message, key ...string) (partition int32, offset int64, err error) {
	kMsg := &sarama.ProducerMessage{}
	kMsg.Topic = p.topic
	if len(key) == 1 {
		kMsg.Key = sarama.StringEncoder(key[0])
	}
	bMsg, err := proto.Marshal(m)
	if err != nil {
		return -1, -1, err
	}
	kMsg.Value = sarama.ByteEncoder(bMsg)
	return p.producer.SendMessage(kMsg)
}
