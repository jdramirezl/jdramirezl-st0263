package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

var (
	robin bool = true
)

func GetVar(r *http.Request, key string) (value string) {
	return r.URL.Query().Get(key)
}

func NewRouter(conn *grpc.ClientConn) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		// GRPC
		if robin {
			log.Printf("Response from GRPC list")
			response, err := List(conn)
			if err != nil {
				log.Fatalf("Error when calling func: %s", err)
			}

			log.Printf("Response from server list: %s", response)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))

			// RabbitMQ
		} else {
			log.Printf("Response from server list: %s", "RabbitMQ")
		}

		// Change the value of robin
		robin = !robin
	})

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// GRPC
		if robin {
			log.Printf("Response from GRPC search")
			name := GetVar(r, "name")
			response, err := Search(conn, name)
			if err != nil {
				log.Fatalf("Error when calling func: %s", err)
			}

			log.Printf("Response from server list: %s", response)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))

			// RabbitMQ
		} else {
			log.Printf("Response from server search: %s", "RabbitMQ")
		}

		
	})

	return mux
}

func runHttp(listenAddr string, conn *grpc.ClientConn) error {
	s := http.Server{
		Addr:    listenAddr,
		Handler: NewRouter(conn), // Our own instance of servemux
	}
	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)
	return s.ListenAndServe()
}

func start() *Configuration {
	config, err := loadConfig("../config")
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return nil
	}
	return config
}

func main() {
	config := start()

	flag.Parse()
	apiAddr := net.JoinHostPort(config.APIIP, config.APIPort)
	grpcAddr := net.JoinHostPort(config.GRPCIP, config.GRPCPort)

	conn, err := CreateConnection(grpcAddr)

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	if err := runHttp(apiAddr, conn); err != nil {
		log.Fatal(err)
	}

}
