package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	const rmqdsn = "amqp://guest:guest@localhost:5672/"

	fmt.Println("Starting Peril server...")

	conn, err := amqp.Dial(rmqdsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	log.Println("connected to", rmqdsn, !conn.IsClosed())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals

	log.Println("graceful shutdown")
}
