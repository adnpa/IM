package transfer

import (
	"github.com/IBM/sarama"
	"github.com/adnpa/IM/common/config"
	"github.com/adnpa/IM/common/constant"
	"github.com/adnpa/IM/common/db/mysql/dao"
	"github.com/adnpa/IM/common/kafka"
	"github.com/adnpa/IM/pkg/pb"
	"google.golang.org/protobuf/proto"
	"log"
)

type PersistentConsumerHandler struct {
	msgHandle               map[string]fcb
	persistentConsumerGroup *kafka.ConsumerGroup
}

func (pc *PersistentConsumerHandler) Init() {
	pc.msgHandle = make(map[string]fcb)

	cfg := &kafka.ConsumerGroupConfig{
		KafkaVersion:   sarama.V0_10_0_0,
		OffsetsInitial: sarama.OffsetNewest,
		IsReturnErr:    false,
	}
	pc.persistentConsumerGroup = kafka.NewConsumerGroup(cfg, []string{config.Config.Kafka.Ws2mschat.Topic}, config.Config.Kafka.Ws2mschat.Addr, config.Config.Kafka.ConsumerGroupID.MsgToMySql)
}

func (pc *PersistentConsumerHandler) handleChatWs2Mysql(msg []byte, msgKey string) {
	msgFromMQ := &pb.MsgDataToMQ{}
	err := proto.Unmarshal(msg, msgFromMQ)
	if err != nil {
		return
	}

	if constant.SingleChatType == msgFromMQ.MsgData.SessionType && msgKey == msgFromMQ.MsgData.RecvID {
		err := dao.InsertChatLog(msgFromMQ)
		if err != nil {
			log.Println(err)
			return
		}
	} else if constant.GroupChatType == msgFromMQ.MsgData.SessionType {
		err := dao.InsertChatLog(msgFromMQ)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (pc *PersistentConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *PersistentConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *PersistentConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Println("kafka get info to mysql", "", "msgTopic", msg.Topic, "msgPartition", msg.Partition, "msg", string(msg.Value))
			pc.msgHandle[msg.Topic](msg.Value, string(msg.Key))
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
