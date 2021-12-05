package consumer

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func Consume() {
	var rabbitHost = os.Getenv("RABBIT_HOST")
	var rabbitPort = os.Getenv("RABBIT_PORT")
	var rabbitUser = os.Getenv("RABBIT_USERNAME")
	var rabbitPassword = os.Getenv("RABBIT_PASSWORD")

	conn, err := amqp.Dial("amqp://" + rabbitUser + ":" + rabbitPassword + "@" + rabbitHost + ":" + rabbitPort + "/")
	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"test-queue",
		false,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to declare a queue")

	fmt.Println("Channel and Queue established")

	msgs, err := ch.Consume(q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to register consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s\n", d.Body)
			d.Ack(false)
		}
	}()

	fmt.Println("Running...")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
