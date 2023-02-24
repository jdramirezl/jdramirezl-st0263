package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Configuration struct {
	Host string
	Port string
    ApiPort   string
    RabbitPort string
    GrpcPort string
}

func loadConfig(directory string) (*Configuration, error) {
    file, err := os.Open(directory + "/.env")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    config := &Configuration{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            continue
        }

        switch parts[0] {
        case "API_PORT":
            config.ApiPort = parts[1]
        case "RABBIT_PORT":
            config.RabbitPort = parts[1]
        case "GRPC_PORT":
            config.GrpcPort = parts[1]
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return config, nil
}

func main() {
    config, err := loadConfig("./my-directory")
    if err != nil {
        fmt.Println("Error loading configuration:", err)
        return
    }

    fmt.Println("API Port:", config.ApiPort)
    fmt.Println("Rabbit Port:", config.RabbitPort)
    fmt.Println("gRPC Port:", config.GrpcPort)
}

