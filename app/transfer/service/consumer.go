package service

import (
	"log"

	"github.com/adnpa/IM/api/pb"
	"github.com/adnpa/IM/pkg/logger"
	amqp "github.com/rabbitmq/amqp091-go"
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

// func TestSend() {
// 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()
// 	go NewConsumer(conn).Run()
// }

func (c *Consumer) Run() {
	ch, err := c.conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
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

func (c *Consumer) handleMsg(msg []byte) error {
	log.Printf("Received a message: %s", msg)
	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		logger.Panicf("%s: %s", msg, err)
	}
}
