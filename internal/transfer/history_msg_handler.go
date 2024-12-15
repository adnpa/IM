package transfer

import (
	"github.com/IBM/sarama"
	"github.com/adnpa/IM/pkg/common/config"
	"github.com/adnpa/IM/pkg/common/db/mongodb"
	"github.com/adnpa/IM/pkg/common/db/redis"
	"github.com/adnpa/IM/pkg/common/kafka"
	"github.com/adnpa/IM/pkg/pb"
	"log"
)

type HistoryConsumerHandler struct {
	msgHandle            map[string]fcb
	historyConsumerGroup *kafka.ConsumerGroup
	singleMsgCount       uint64
	groupMsgCount        uint64
}

func (hc *HistoryConsumerHandler) Init() {
	//todo 统计 statistics
	hc.msgHandle = make(map[string]fcb)
	hc.singleMsgCount = 0
	hc.groupMsgCount = 0

	cfg := &kafka.ConsumerGroupConfig{
		KafkaVersion:   sarama.V0_10_0_0,
		OffsetsInitial: sarama.OffsetNewest,
		IsReturnErr:    false,
	}
	hc.historyConsumerGroup = kafka.NewConsumerGroup(cfg, []string{config.Config.Kafka.Ws2mschat.Topic}, config.Config.Kafka.Ws2mschat.Addr, config.Config.Kafka.ConsumerGroupID.MsgToMongo)
}

func (hc *HistoryConsumerHandler) handleChatWs2Mongo(msg []byte, msgKey string) {

}

func saveUserChat(uid string, msg *pb.MsgDataToMQ) error {
	//time := utils.NowMilliSecond()
	seq, err := redis.IncrUserSeq(uid)
	if err != nil {
		return err
	}
	msg.MsgData.Seq = uint32(seq)
	pbSaveData := pb.MsgDataToDB{}
	pbSaveData.MsgData = msg.MsgData
	return mongodb.SaveUserChatMongo2(uid, pbSaveData.MsgData.SendTime, &pbSaveData)
}

func (hc *HistoryConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (hc *HistoryConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (hc *HistoryConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Println("kafka get info to mongo", "", "msgTopic", msg.Topic, "msgPartition", msg.Partition, "msg", string(msg.Value))
			hc.msgHandle[msg.Topic](msg.Value, string(msg.Key))
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
