package publisher

import (
	"fmt"
	"github.com/NielsBattle/GO-RabbitMQ/publisher/pkg/helper"
	"github.com/julienschmidt/httprouter"
	"github.com/streadway/amqp"
	"os"
)

func Publish(params httprouter.Params) {
	var rabbitHost = os.Getenv("RABBIT_HOST")
	var rabbitPort = os.Getenv("RABBIT_PORT")
	var rabbitUser = os.Getenv("RABBIT_USERNAME")
	var rabbitPassword = os.Getenv("RABBIT_PASSWORD")

	conn, err := amqp.Dial("amqp://" + rabbitUser + ":" + rabbitPassword + "@" + rabbitHost + ":" + rabbitPort + "/")
	helper.FailOnError(err, "Failed to connect to RabbitMQ")

	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			helper.FailOnError(err, "Failed to close connection!")
		}
	}(conn)

	ch, err := conn.Channel()
	helper.FailOnError(err, "Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			helper.FailOnError(err, "Failed to close channel!")
		}
	}(ch)


	fmt.Println("Received message to publish on Queue!")

	q, err := ch.QueueDeclare(
		"test-queue",
		true,
		false,
		false,
		false,
		nil)
	helper.FailOnError(err, "Failed to declare a queue")

	message := params.ByName("message")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plaing",
			Body: []byte(message),
		})
	helper.FailOnError(err, "Failed to publish a message")


	//errCh := ch.Close()
	//if errCh != nil {
	//	fmt.Println("Failed to close channel.")
	//}
	//errConn := conn.Close()
	//if errConn != nil {
	//	fmt.Println("Failed to close connection." + err.Error())
	//	return
	//}
	//return
}
