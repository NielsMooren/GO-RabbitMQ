package main

import (
	"fmt"
	"github.com/NielsBattle/GO-RabbitMQ/publisher/pkg/publisher"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {

	router := httprouter.New()

	router.POST("/publish/:message", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		go publisher.Publish(params)
		writer.WriteHeader(200)
		writer.Write([]byte("hallo"))
	})

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	fmt.Println("Publisher started")
	log.Fatalln(http.ListenAndServe(":8082", router))
}