package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewConsumer(c *amqp.Connection) *Consumer {
	return &Consumer{
		conn: c,
	}
}

type HandleCallback func()

type Consumer struct {
	conn *amqp.Connection
	// Topic         string
	// PartitionList []int32
	// Consumer      sarama.Consumer
	// WG            sync.WaitGroup
	// handleFunc *
}

func TestSend() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	go NewConsumer(conn).Run()
}

func (c *Consumer) Run() {
	ch, err := c.conn.Channel()
	failOnError(err, "Failed to open a channel")
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

	// var forever chan struct{}

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		// todo handle msg
		// chat.MyServer.HandleMsg()
		d.Ack(false)
	}
	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// <-forever
}

// func NewKafkaConsumer(addr []string, topic string) *Consumer {
// 	consumer, err := sarama.NewConsumer(addr, nil)
// 	if err != nil {
// 		panic(err)
// 		return nil
// 	}

// 	partitionList, err := consumer.Partitions(topic)
// 	if err != nil {
// 		panic(err)
// 		return nil
// 	}

// 	p := &Consumer{
// 		addr:          addr,
// 		Topic:         topic,
// 		PartitionList: partitionList,
// 		Consumer:      consumer,
// 	}
// 	return p
// }
