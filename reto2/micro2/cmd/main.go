package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"time"

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
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	fmt.Println("Waiting for messages")

	var forever chan struct{}

	go func() {
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
				fmt.Println("Searching for file: ", req.Args[0])
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
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func bootstrap() *Configuration {
	fmt.Println("Loading configuration")
	config, err := loadConfig("./config")
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

	time.Sleep(5 * time.Second)
	// Run RabbitMQ
	// connRabbit := runRabbitMQ(addr)
	fmt.Println("Startooo")
	fmt.Println("Connecting to RabbitMQ server at", addr)
	connRabbit, err := amqp.Dial("amqp://guest:guest@" + addr + "/")
	failOnError(err, "Failed to connect to RabbitMQ")

	fmt.Println("Connected successfully")

	time.Sleep(5 * time.Second)
	// Create channel
	// chRabbit := CreateChannel(connRabbit)
	chRabbit, err := connRabbit.Channel()
	failOnError(err, "Failed to open a channel")
	fmt.Println("Created channel!")

	// Create Queue
	qName := config.QName
	fmt.Println("Creating queue", qName)
	qRabbit := CreateQueue(chRabbit, qName)

	// Receive messages
	receive(qRabbit, chRabbit)
	defer connRabbit.Close()
}
