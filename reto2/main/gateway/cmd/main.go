package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func bootstrap() *Configuration {
	fmt.Println("Loading configuration...")
	config, err := loadConfig("./config")
	failOnError(err, "Failed to load configuration")
	fmt.Println("Configuration loaded successfully")
	return config
}

func main() {
	// Start
	fmt.Println("Starting THE gateway...")
	config := bootstrap()

	// Create addresses
	apiAddr := ":" + config.APIPort //net.JoinHostPort(config.APIIP, config.APIPort)
	grpcAddr := net.JoinHostPort(config.GRPCIP, config.GRPCPort)
	rabbitAddr := net.JoinHostPort(config.RabbitIP, config.RabbitPort)

	// Run GRPC
	connGRPC := runGRPC(grpcAddr)
	defer connGRPC.Close()

	// Run RabbitMQ
	time.Sleep(5 * time.Second)
	connRabbit := runRabbitMQ(rabbitAddr)
	defer connRabbit.Close()

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
