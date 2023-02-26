package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ReqRabbit struct {
	Type string   `json:"action"`
	Args []string `json:"args,omitempty"`
}

type ResRabbit struct {
	Files []string `json:"files"`
}

func joinResponseFiles(resp ResRabbit) string {
	var result string
	for i, file := range resp.Files {
		if i > 0 {
			result += ","
		}
		result += file
	}
	return result
}

func StructToJSON(message interface{}) ([]byte, error) {
	return json.Marshal(message)
}

func JSONToStruct(jsonData []byte, message interface{}) (interface{}, error) {
	err := json.Unmarshal(jsonData, &message)
	return message, err
}

func send(ch *amqp.Channel, q amqp.Queue, req ReqRabbit) (ResRabbit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ID for the request
	corrId := uuid.New().String()

	// Response queue
	replyQueue := CreateQueue(ch, "")

	// Consumer for res
	msgs, err := ch.Consume(
		replyQueue.Name, // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	failOnError(err, "Failed to register a consumer")

	jsonData, err := StructToJSON(req)
	failOnError(err, "Failed to convert request to JSON")

	// Publish with id
	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			Body:          jsonData,
			ReplyTo:       replyQueue.Name,
			CorrelationId: corrId,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", jsonData)

	// wait for the response message
	for msg := range msgs {
		if msg.CorrelationId == corrId {
			var res ResRabbit
			JSONToStruct(msg.Body, res)
			failOnError(err, "Failed to convert JSON to struct")
			return res, nil
		}
	}

	return ResRabbit{}, fmt.Errorf("timed out waiting for a response")
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

func runRabbitMQ(listenAddr string) *amqp.Connection {
	// RabbitMQ
	fmt.Println("Connecting to RabbitMQ server at", listenAddr)
	connRabbit, err := amqp.Dial("amqp://guest:guest@" + listenAddr + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connRabbit.Close()
	fmt.Println("Connected successfully")
	return connRabbit
}
