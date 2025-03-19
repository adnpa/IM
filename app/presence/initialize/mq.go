package initialize

import (
	"github.com/adnpa/IM/app/presence/global"
	"github.com/adnpa/IM/pkg/common/mq/rabbitmq"
	"github.com/adnpa/IM/pkg/logger"
	"go.uber.org/zap"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitProducer() {
	conn, err := amqp.Dial("amqp://admin:passwd@localhost:5672/")
	if err != nil {
		logger.Panic("Failed to connect to RabbitMQ", zap.Error(err))
	}
	defer conn.Close()
	global.Producer = rabbitmq.NewProducer(conn, "transfer")
}
