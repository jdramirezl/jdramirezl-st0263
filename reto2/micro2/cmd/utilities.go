package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func runRabbitMQ(listenAddr string) *amqp.Connection {
	// RabbitMQ
	fmt.Println("Connecting to RabbitMQ server at", listenAddr)
	connRabbit, err := amqp.Dial("amqp://guest:guest@" + listenAddr + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connRabbit.Close()
	fmt.Println("Connected successfully")
	return connRabbit
}

func CreateQueue(ch *amqp.Channel, name string) amqp.Queue {
	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return q
}

func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch
}

