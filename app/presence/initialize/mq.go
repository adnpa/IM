package initialize

import (
	"fmt"

	"github.com/adnpa/IM/app/presence/global"
	"github.com/adnpa/IM/pkg/common/mq/rabbitmq"
	"github.com/adnpa/IM/pkg/logger"
	"go.uber.org/zap"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitProducer() {
	rabbitCfg := global.ServerConfig.RabbitMQInfo
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", rabbitCfg.User, rabbitCfg.Password, rabbitCfg.Host, rabbitCfg.Port))
	if err != nil {
		logger.Panic("Failed to connect to RabbitMQ", zap.Error(err))
	}
	// defer conn.Close()
	global.Producer = rabbitmq.NewProducer(conn, "transfer")
}
