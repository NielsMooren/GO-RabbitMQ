package consumer

import (
	"context"
	"fmt"
	"github.com/Azure/go-amqp"
	"log"
	"os"
)

var rabbitHost = os.Getenv("RABBIT_HOST")
var rabbitport = os.Getenv("RABBIT_PORT")
var rabbitUser = os.Getenv("RABBIT_USERNAME")
var rabbitpassword = os.Getenv("RABBIT_PASSWORD")

func consume() {
	conn, err := amqp.Dial("amqps://" + rabbitUser + ":" + rabbitpassword + "@" + rabbitHost + "/")

	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ: "+rabbitHost, err)
	}

	ctx := context.Background()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to create new AMQP session "+rabbitHost, err)
	}

	queue, err := session.NewReceiver(amqp.LinkSourceAddress("/test-queue"), amqp.LinkCredit(10))
	if err != nil {
		log.Fatalf("Creating listener failed " + rabbitHost)
	}

	msg, err := queue.Receive(ctx)
	if err != nil {
		log.Fatalf("Failed to receive message")
	}

	//TODO test the ability to receive messages from queue

	fmt.Println(msg.Data)
	fmt.Println(msg.Header)
	fmt.Println(msg.Format)
	fmt.Println(msg.Value)

}
