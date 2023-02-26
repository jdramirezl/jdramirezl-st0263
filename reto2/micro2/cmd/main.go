package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	dir string
)

type ReqRabbit struct {
	Type string   `json:"action"`
	Args []string `json:"args,omitempty"`
}

type ResRabbit struct {
	Files []string `json:"files"`
}

func list() (ResRabbit, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return ResRabbit{}, err
	}

	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileNames = append(fileNames, file.Name())
	}

	res := ResRabbit{Files: fileNames}

	return res, nil
}

func search(keyword string) (ResRabbit, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return ResRabbit{}, err
	}

	var found bool = false
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name() == keyword {
			found = true
			break
		}
	}

	var res ResRabbit
	if found {
		res = ResRabbit{Files: []string{"Found"}}
	} else {
		res = ResRabbit{Files: []string{}}
	}

	return res, nil
}

func receive(q amqp.Queue, ch *amqp.Channel) {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// Listen for incoming requests
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)

		// Parse the request
		var req ReqRabbit
		err := json.Unmarshal(d.Body, &req)
		failOnError(err, "Failed to parse request")

		// Process the request and generate the response
		var res ResRabbit
		switch req.Type {
		case "list":
			res, err = list()
			failOnError(err, "Failed to list files")
		case "search":
			res, err = search(req.Args[0])
			failOnError(err, "Failed to search for file")
		default:
			failOnError(fmt.Errorf("Unknown action type: %s", req.Type), "")
		}

		// Convert the response to JSON
		jsonData, err := json.Marshal(res)
		failOnError(err, "Failed to convert response to JSON")

		// Declare a new response queue
		resQueue, err := ch.QueueDeclare(
			"",    // name
			false, // durable
			true,  // delete when unused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")

		// Send the response back to the appropriate queue
		err = ch.Publish(
			"",        // exchange
			d.ReplyTo, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType:   "application/json",
				CorrelationId: d.CorrelationId,
				Body:          jsonData,
				ReplyTo:       resQueue.Name,
			},
		)
		failOnError(err, "Failed to publish a message")

		log.Printf("Response sent: %s", jsonData)
	}
}

func bootstrap() *Configuration {
	fmt.Println("Loading configuration")
	config, err := loadConfig("../config")
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return nil
	}
	fmt.Println("Configuration loaded succesfully")
	return config
}

func main() {
	fmt.Println("Starting RabbitMQ server")
	config := bootstrap()
	dir = config.Directory

	addr := net.JoinHostPort(config.IP, config.Port)

	// Run RabbitMQ
	connRabbit := runRabbitMQ(addr)

	// Create channel
	chRabbit := CreateChannel(connRabbit)

	// Create Queue
	qName := config.QName
	qRabbit := CreateQueue(chRabbit, qName)

	// Receive messages
	receive(qRabbit, chRabbit)
}
