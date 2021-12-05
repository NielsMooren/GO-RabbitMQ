package main

import (
	"fmt"
	"github.com/NielsBattle/GO-RabbitMQ/pkg/publisher"
	"github.com/julienschmidt/httprouter"
	"net/http"

	//"github.com/NielsBattle/GO-RabbitMQ/pkg/publisher"
	"github.com/joho/godotenv"
	//"github.com/julienschmidt/httprouter"
	"log"
	//"net/http"
)

func main() {
	router := httprouter.New()

	router.POST("/publish/:message", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		publisher.Publish(writer, request, params)
	})


	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	fmt.Println("Hello world")
	//publisher.Publish()
	//consumer.Consume()
	log.Fatalln(http.ListenAndServe(":8080", router))
}
