package rabbitmq

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	conn *amqp.Connection

	// topic    string
	// addr     []string
	// config   *sarama.Config
	// producer sarama.SyncProducer
}

func NewProducer(conn *amqp.Connection, topic string) *Producer {
	return &Producer{
		conn: conn,
	}
}

func (c *Producer) Send(name string, uid string, data []byte) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	headers := make(amqp.Table)
	headers["uid"] = uid
	return ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Headers:     headers,
			Body:        data,
		})
}
