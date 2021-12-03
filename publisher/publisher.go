package publisher

import (
	"context"
	"github.com/Azure/go-amqp"
	"log"
	"os"
	"time"
)

//TODO add .env file
var rabbitHost = os.Getenv("RABBIT_HOST")
var rabbitport = os.Getenv("RABBIT_PORT")
var rabbitUser = os.Getenv("RABBIT_USERNAME")
var rabbitpassword = os.Getenv("RABBIT_PASSWORD")

func publish() {
	conn, err := amqp.Dial("amqps://" + rabbitUser + ":" + rabbitpassword + "@" + rabbitHost + "/")

	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ: "+rabbitHost, err)
	}

	ctx := context.Background()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to create new AMQP session "+rabbitHost, err)
	}

	sender, err := session.NewSender(amqp.LinkTargetAddress("/test-queue"))
	if err != nil {
		log.Fatalf("Failed to create new publisher for host: " + rabbitHost)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	if err != nil {
		log.Fatalf("Failed to push message on queue with host: " + rabbitHost)
	}

	sender.Close(ctx)
	cancel()

	// TODO fix sending option

	//TODO test ability to push messages onto queue.

	//queue.Send(ctx, amqp.NewMessage())

}
