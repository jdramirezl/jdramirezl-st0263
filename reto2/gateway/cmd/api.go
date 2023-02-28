package main

import (
	"fmt"
	"log"
	"net/http"
	m1 "retos/reto2/gateway/internal"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

var (
	robin bool = true
)

func output(w http.ResponseWriter, response string, fail string, err error, len int) {
	// Error handling
	if err != nil {
		log.Fatalf("Error when calling func: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		log.Printf("Response from server list: %s", response)
		w.WriteHeader(http.StatusOK)

		if len == 0 {
			w.Write([]byte(fail))
		} else {
			w.Write([]byte(response))
		}
	}
}

func GetVar(r *http.Request, key string) (value string) {
	return r.URL.Query().Get(key)
}

func NewRouter(conn *grpc.ClientConn, chanRabbit *amqp.Channel, qRabbit amqp.Queue) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		var (
			files    []string
			response string
			err      error
		)

		if robin {
			// GRPC
			var resGRPC *m1.FileResponse
			log.Printf("Response from GRPC: list")
			resGRPC, err = List(conn)
			failOnError(err, "Failed to call List")
			files = resGRPC.Name
		} else {
			// RabbitMQ
			var resRabbit ResRabbit
			log.Printf("Response from RabbitMQ: list")
			req := ReqRabbit{Type: "list", Args: []string{}}
			resRabbit, err = send(chanRabbit, qRabbit, req)
			failOnError(err, "Failed to call List")
			fmt.Printf("Response from RabbitMQ: %s", resRabbit.Files)
			files = resRabbit.Files
		}

		// Join the files
		response = strings.Join(files, ", ")

		// Output
		output(w, response, "No files found", err, len(response))

		// Change the value of robin
		robin = !robin
	})

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		var (
			err      error
			name     string
			response string
		)

		name = GetVar(r, "name")

		if robin {
			// GRPC
			log.Printf("Response from GRPC: search")
			var resGRPC *m1.FileResponse
			resGRPC, err = Search(conn, name)

			if len(resGRPC.Name) == 0 {
				response = ""
			} else {
				response = resGRPC.Name[0]
			}

		} else {
			// RabbitMQ
			log.Printf("Response from RabbitMQ: search")
			var resRabbit ResRabbit
			req := ReqRabbit{Type: "search", Args: []string{name}}
			resRabbit, err = send(chanRabbit, qRabbit, req)

			if len(resRabbit.Files) == 0 {
				response = ""
			} else {
				response = resRabbit.Files[0]
			}
		}

		// Output
		output(w, response, "No files found", err, len(response))

		// Change the value of robin
		robin = !robin
	})

	return mux
}

func RunHttp(listenAddr string, connGRPC *grpc.ClientConn, chanRabbit *amqp.Channel, qRabbit amqp.Queue) error {
	s := http.Server{
		Addr:    listenAddr,
		Handler: NewRouter(connGRPC, chanRabbit, qRabbit), // Our own instance of servemux
	}
	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)
	return s.ListenAndServe()
}
