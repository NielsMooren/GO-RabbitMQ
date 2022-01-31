package main

import (
	"fmt"
	"github.com/NielsBattle/GO-RabbitMQ/consumer/pkg/consumer"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {

	router := httprouter.New()

	router.POST("/consume", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		consumer.Consume(writer)
	})

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	fmt.Println("Consumer started!")
	log.Fatalln(http.ListenAndServe(":8081", router))
}
