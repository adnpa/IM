package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/app/transfer/global"
	"github.com/adnpa/IM/internal/model"
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

	errChan := make(chan *amqp.Error)

	q, err := ch.QueueDeclare(
		"im_message", // name
		true,         // durable
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

	for {
		select {
		case d, ok := <-msgs:
			if !ok {
				logger.Info("Message channel closed")
				return
			}
			err := c.handleMsg(d.Body)
			if err != nil {
				logger.Error("handle msg fail, back to mq", zap.Error(err))
				d.Nack(false, false)
			} else {
				d.Ack(false)
			}
			// time.Sleep(5 * time.Second)
		case err := <-errChan:
			if err != nil {
				logger.Error("Channel closed", zap.Error(err))
			}
			return
		}
	}
}

func (c *Consumer) handleMsg(data []byte) error {
	msg := &model.CommonMsg{}
	err := json.Unmarshal(data, msg)
	if err != nil {
		logger.Warn("unmarshal fail", zap.Error(err))
		return err
	}
	logger.Info("consume msg from mq", zap.Any("msg", msg))

	switch msg.Cmd {
	case model.TypSingle:
		return c.handleSingle(&msg.ChatMsg)
	case model.TypGroup:
		// return c.handleGroup(msg)
	case model.TypMsgAckFromClient:
		msg.Cmd = model.TypMsgAckFromServerForRecver
		_, err = global.PresenceCli.SendMsg(context.Background(),
			&pb.SendMsgReq{UserId: int32(msg.ChatMsg.From),
				Msg: &pb.ChatMsg{Typ: int32(model.TypMsgAckFromServerForRecver), Id: msg.AckMsg.Id, Seq: int64(msg.AckMsg.Seq)}})
		if err != nil {
			logger.Warn("call presence service fail", zap.Error(err))
			return err
		}
	default:
		logger.Warn("unrecognized cmd")
	}
	return nil
}

func (c *Consumer) handleSingle(msg *model.Message) error {
	if msg.To == 0 || msg.From == 0 || msg.To == msg.From {
		logger.Warn("error format", zap.Any("msg", msg))
		return fmt.Errorf("bad msg format")
	}

	// TODO: 序列号去重

	// TODO: 分布式id服务
	msg.Id = utils.NowMilliSecond()

	pbMsg := &pb.ChatMsg{Typ: int32(msg.Cmd)}
	copier.Copy(pbMsg, msg)

	resp, err := global.PresenceCli.IsOnline(context.Background(), &pb.IsOnlineReq{UserId: int32(msg.To)})
	if err != nil {
		logger.Warn("call presence service fail", zap.Error(err))
		return err
	}

	// 总是先入库 收到用户ack再把数据删掉
	putMsgResp, err := global.OffineCli.PutMsg(context.Background(), &pb.PutMsgReq{UserId: int32(msg.To), Msg: pbMsg})
	if err != nil || !putMsgResp.Succ {
		logger.Warn("call offline service fail", zap.Error(err))
		return err
	}

	// 回复ack
	_, err = global.PresenceCli.SendMsg(context.Background(),
		&pb.SendMsgReq{UserId: int32(msg.From),
			Msg: &pb.ChatMsg{Typ: int32(model.TypMsgAckFromServerForSender), Id: msg.Id, Seq: msg.Seq}})
	if err != nil {
		logger.Warn("call presence service fail", zap.Error(err))
		return err
	}

	if resp.IsOnline {
		logger.Info("user online, send")
		_, err := global.PresenceCli.SendMsg(context.Background(), &pb.SendMsgReq{UserId: int32(msg.To), Msg: pbMsg})
		if err != nil {
			logger.Warn("call presence service fail", zap.Error(err))
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
