package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	m1 "retos/reto2/gateway/internal"
	"strings"

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
		var (
			response *m1.FileResponse
			err      error
		)

		if robin {
			// GRPC
			log.Printf("Response from GRPC: list")
			response, err = List(conn)
		} else {
			// RabbitMQ
			log.Printf("Response from RabbitMQ: list")
		}

		// Error handling
		if err != nil {
			log.Fatalf("Error when calling func: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			log.Printf("Response from server list: %s", response)
			w.WriteHeader(http.StatusOK)

			if len(response.Name) == 0 {
				w.Write([]byte("No files found"))
			} else {
				w.Write([]byte(strings.Join(response.Name, ", ")))
			}
		}

		// Change the value of robin
		robin = !robin
	})

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		var (
			response *m1.FileResponse
			err      error
		)

		if robin {
			// GRPC
			log.Printf("Response from GRPC: search")
			name := GetVar(r, "name")
			response, err = Search(conn, name)
		} else {
			// RabbitMQ
			log.Printf("Response from RabbitMQ: search")
		}

		// Error handling
		if err != nil {
			log.Fatalf("Error when calling func: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			log.Printf("Response from server list: %s", response)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.Name[0]))
		}

		// Change the value of robin
		robin = !robin

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
	fmt.Println("Loading configuration...")
	config, err := loadConfig("../config")
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return nil
	}
	fmt.Println("Configuration loaded successfully")
	return config
}

func main() {
	fmt.Println("Starting gateway...")
	config := start()

	flag.Parse()
	apiAddr := net.JoinHostPort(config.APIIP, config.APIPort)
	grpcAddr := net.JoinHostPort(config.GRPCIP, config.GRPCPort)

	conn, err := CreateConnection(grpcAddr)

	fmt.Println("Connecting to GRPC server at", grpcAddr)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	fmt.Println("Connected successfully")

	if err := runHttp(apiAddr, conn); err != nil {
		log.Fatal(err)
	}

}
