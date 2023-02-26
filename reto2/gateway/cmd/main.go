package main

import (
	"fmt"
	"log"
	"net"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func bootstrap() *Configuration {
	fmt.Println("Loading configuration...")
	config, err := loadConfig("../config")
	failOnError(err, "Failed to load configuration")
	fmt.Println("Configuration loaded successfully")
	return config
}

func main() {
	// Start
	fmt.Println("Starting gateway...")
	config := bootstrap()

	// Create addresses
	apiAddr := net.JoinHostPort(config.APIIP, config.APIPort)
	grpcAddr := net.JoinHostPort(config.GRPCIP, config.GRPCPort)
	rabbitAddr := net.JoinHostPort(config.RabbitIP, config.RabbitPort)

	// Run GRPC
	connGRPC := runGRPC(grpcAddr)

	// Run RabbitMQ
	connRabbit := runRabbitMQ(rabbitAddr)

	// Create channel
	chRabbit := CreateChannel(connRabbit)

	// Create Queue
	qName := config.RabbitQ
	qRabbit := CreateQueue(chRabbit, qName)

	// Run HTTP
	if err := RunHttp(apiAddr, connGRPC, chRabbit, qRabbit); err != nil {
		log.Fatal(err)
	}
}
