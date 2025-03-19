package tests

import (
	"context"
	"log"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func TestSendMsg(t *testing.T) {
	sendMsg := "Hello World!"

	conn, err := amqp.Dial("amqp://admin:password@localhost:5672/")
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Error(err)
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
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(sendMsg),
		})
	if err != nil {
		t.Error(err)
	}
	log.Printf(" [x] Sent %s\n", sendMsg)
}
