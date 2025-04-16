package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

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
			uid := d.Headers["uid"].(string)
			err := c.handleMsg(uid, d.Body)
			if err != nil {
				logger.Error("handle msg fail, back to mq", zap.Error(err))
				d.Nack(false, false)
			} else {
				d.Ack(false)
				intUid, _ := strconv.ParseInt(uid, 10, 64)
				global.PresenceCli.SendMsg(context.Background(),
					&pb.SendMsgReq{UserId: int32(intUid), Type: int32(model.CmdAck)})
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

func (c *Consumer) handleMsg(sourceId string, data []byte) error {
	msg := &model.CommonMsg{}
	err := json.Unmarshal(data, msg)
	if err != nil {
		logger.Warn("unmarshal fail", zap.Error(err))
		return err
	}
	logger.Info("consume msg from mq", zap.Any("msg", msg))

	switch msg.Cmd {
	case model.CmdChat:
		var chatMsg model.ChatMessage
		err = json.Unmarshal(msg.Data, &chatMsg)
		if err != nil {
			logger.Warn("msg body format error", zap.Error(err), zap.Any("data", msg.Data))
			log.Fatal(err)
		}
		switch chatMsg.Type {
		case model.ChatTypeSingle:
			return c.handleSingle(&chatMsg)
		case model.ChatTypeGroup:
			return c.handleGroup(&chatMsg)
		default:
			logger.Warn("unreconigze chat type")
		}
	// case model.TypMsgAckFromClient:
	// 	msg.Cmd = model.TypMsgAckFromServerForRecver
	// 	_, err = global.PresenceCli.SendMsg(context.Background(),
	// 		&pb.SendMsgReq{UserId: int32(msg.ChatMsg.From), Type: int32(model.TypMsgAckFromServerForRecver),
	// 			Msg: &pb.ChatMsg{Typ: int32(model.TypMsgAckFromServerForRecver), Id: msg.AckMsg.Id, Seq: int64(msg.AckMsg.Seq)}})
	// 	if err != nil {
	// 		logger.Warn("call presence service fail", zap.Error(err))
	// 		return err
	// 	}
	case model.CmdPullOfflineMsgs:
		err = c.handlePullOffline(sourceId)
		if err != nil {
			logger.Warn("pull offline msg fail", zap.Error(err))
			return err
		}
	// case model.TypAckPullOffline:
	// 	global.OffineCli.RemoveMsg(context.Background(), &pb.RemoveMsgReq{Uid: int32(msg.PullOfflineRespAck.Uid)})
	default:
		logger.Warn("unrecognized cmd", zap.Any("cmd", msg.Cmd))
	}
	return nil
}

func (c *Consumer) handlePullOffline(uid string) error {
	intUid, err := strconv.ParseInt(uid, 10, 32)
	if err != nil {
		return err
	}
	resp, err := global.OffineCli.GetOfflineMsg(context.Background(), &pb.GetOfflineMsgReq{Uid: int32(intUid)})
	if err != nil {
		return err
	}

	var msgs []model.ChatMessage
	for _, pbMsg := range resp.Msgs {
		msg := model.ChatMessage{}
		copier.Copy(&msg, pbMsg)
		msgs = append(msgs, msg)
	}
	data, err := json.Marshal(msgs)
	if err != nil {
		return err
	}
	_, err = global.PresenceCli.SendMsg(context.Background(), &pb.SendMsgReq{UserId: int32(intUid),
		Type: int32(model.CmdPullOfflineMsgs), Body: data})
	return err
}

func (c *Consumer) handleSingle(msg *model.ChatMessage) error {
	if msg.To == 0 || msg.From == 0 || msg.To == msg.From {
		logger.Warn("error format", zap.Any("msg", msg))
		return fmt.Errorf("bad msg format")
	}

	// TODO: 序列号去重

	// TODO: 分布式id服务
	msg.Id = utils.NowMilliSecond()

	pbMsg := &pb.ChatMsg{}
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

	// TODO:回复ack
	// _, err = global.PresenceCli.SendMsg(context.Background(),
	// 	&pb.SendMsgReq{UserId: int32(msg.From), Type: int32(model.TypMsgAckFromServerForSender),
	// 		Msg: &pb.ChatMsg{Type: int32(model.TypMsgAckFromServerForSender), Id: msg.Id, Seq: msg.Seq}})
	// if err != nil {
	// 	logger.Warn("call presence service fail", zap.Error(err))
	// 	return err
	// }

	if resp.IsOnline {
		logger.Info("user online, send")
		data, err := json.Marshal(msg)
		if err != nil {
			return err
		}
		_, err = global.PresenceCli.SendMsg(context.Background(), &pb.SendMsgReq{UserId: int32(msg.To), Type: int32(msg.Type), Body: data})
		if err != nil {
			logger.Warn("call presence service fail", zap.Error(err))
			return err
		}
	}
	return nil
}

func (c *Consumer) handleGroup(msg *model.ChatMessage) error {
	if msg.To == 0 || msg.From == 0 || msg.To == msg.From {
		logger.Warn("error format", zap.Any("msg", msg))
		return fmt.Errorf("bad msg format")
	}

	msg.Id = utils.NowMilliSecond()

	pbMsg := &pb.ChatMsg{Type: int32(msg.Type)}
	copier.Copy(pbMsg, msg)

	resp, err := global.GroupCli.GetGroupMemberById(context.Background(), &pb.GetGroupMemberByIdReq{GroupId: msg.To})
	if err != nil {
		return err
	}
	// logger.Infof("send to all group members", "members", users)
	// 对群里所有用户 复制一条消息到队列
	for _, member := range resp.Members {
		if member.UserId == pbMsg.From {
			continue
		}

		logger.Info("send to member", zap.Any("user id", member.UserId), zap.Any("msg", pbMsg))

		_, err := global.OffineCli.PutMsg(context.Background(), &pb.PutMsgReq{UserId: member.UserId, Msg: pbMsg})
		if err != nil {
			logger.Error("call offline fail", zap.Error(err))
			return err
		}

		preResp, err := global.PresenceCli.IsOnline(context.Background(), &pb.IsOnlineReq{UserId: member.UserId})
		if err != nil {
			logger.Error("call presence fail", zap.Error(err))
			return err
		}
		if preResp.IsOnline {
			logger.Info("member online", zap.Any("uid", member.UserId))
			data, err := json.Marshal(msg)
			if err != nil {
				return err
			}
			global.PresenceCli.SendMsg(context.Background(), &pb.SendMsgReq{UserId: member.UserId, Type: pbMsg.Type, Body: data})
		}
	}

	// 回复ack
	// _, err = global.PresenceCli.SendMsg(context.Background(),
	// 	&pb.SendMsgReq{UserId: int32(msg.From), Type: int32(model.TypMsgAckFromServerForSender),
	// 		Msg: &pb.ChatMsg{Type: int32(model.TypMsgAckFromServerForSender), Id: msg.Id, Seq: msg.Seq}})
	// if err != nil {
	// 	logger.Warn("call presence service fail", zap.Error(err))
	// 	return err
	// }

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		logger.Panicf("%s: %s", msg, err)
	}
}
