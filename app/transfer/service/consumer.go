package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/transfer/global"
	"github.com/adnpa/IM/app/transfer/model"
	"github.com/adnpa/IM/internal/utils"
	"github.com/adnpa/IM/pkg/logger"
	"github.com/jinzhu/copier"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func NewConsumer(c *amqp.Connection) *Consumer {
	return &Consumer{
		conn: c,
	}
}

type Consumer struct {
	pb.UnimplementedTransferServer

	conn *amqp.Connection
}

func (c *Consumer) Run() {
	ch, err := c.conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"im_message", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {
		err := c.handleMsg(d.Body)
		if err != nil {
			d.Ack(false)
		}
		d.Ack(true)
	}
}

func (c *Consumer) handleMsg(data []byte) error {
	msg := &model.Message{}
	err := json.Unmarshal(data, msg)
	if err != nil {
		return err
	}
	logger.Debug("recv msg", zap.Any("msg", msg))

	switch msg.Cmd {
	case model.TypSingle:
		c.handleSingle(msg)
	case model.TypGroup:
		c.handleGroup(msg)
	default:
	}
	return nil
}

func (c *Consumer) handleSingle(msg *model.Message) error {
	if msg.To == msg.From {
		return fmt.Errorf("send msg to self is not support")
	}

	// TODO: 序列号去重

	// TODO: 分布式id服务
	msg.Id = utils.NowMilliSecond()

	pbMsg := &pb.ChatMsg{}
	copier.Copy(pbMsg, msg)

	resp, err := global.PresenceCli.IsOnline(context.Background(), &pb.IsOnlineReq{UserId: int32(msg.To)})
	if err != nil {
		return err
	}

	// 总是先入库 收到用户ack再把数据删掉
	putMsgResp, err := global.OffineCli.PutMsg(context.Background(), &pb.PutMsgReq{UserId: int32(msg.To), Msg: pbMsg})
	if err != nil || !putMsgResp.Succ {
		return err
	}

	if resp.IsOnline {
		_, err := global.PresenceCli.SendMsg(context.Background(), &pb.SendMsgReq{UserId: int32(msg.To), Msg: pbMsg})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Consumer) handleGroup(msg *model.Message) error {
	msg.Id = utils.NowMilliSecond()

	pbMsg := &pb.ChatMsg{}
	copier.Copy(pbMsg, msg)

	resp, err := global.GroupCli.GetGroupMemberById(context.Background(), &pb.GetGroupMemberByIdReq{GroupId: msg.To})
	if err != nil {
		return err
	}
	// logger.Infof("send to all group members", "members", users)
	// 对群里所有用户 复制一条消息到队列
	for _, member := range resp.Members {
		_, err := global.OffineCli.PutMsg(context.Background(), &pb.PutMsgReq{UserId: member.UserId, Msg: pbMsg})
		if err != nil {
			return err
		}

		preResp, err := global.PresenceCli.IsOnline(context.Background(), &pb.IsOnlineReq{UserId: member.UserId})
		if err != nil {
			return err
		}
		if preResp.IsOnline {
			global.PresenceCli.SendMsg(context.Background(), &pb.SendMsgReq{Msg: pbMsg})
		}
	}
	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		logger.Panicf("%s: %s", msg, err)
	}
}
