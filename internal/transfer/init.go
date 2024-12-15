package transfer

import (
	"github.com/adnpa/IM/common/config"
	"github.com/adnpa/IM/common/kafka"
)

var (
	persistentCH PersistentConsumerHandler
	historyCH    HistoryConsumerHandler
	producer     *kafka.Producer
)

func Init() {
	persistentCH.Init()
	historyCH.Init()
	producer = kafka.NewProducer(config.Config.Kafka.Ms2pschat.Addr, config.Config.Kafka.Ms2pschat.Topic)
}
func Run() {
	//register mysqlConsumerHandler to
	go persistentCH.persistentConsumerGroup.RegisterHandleAndConsumer(&persistentCH)
	go historyCH.historyConsumerGroup.RegisterHandleAndConsumer(&historyCH)
}
