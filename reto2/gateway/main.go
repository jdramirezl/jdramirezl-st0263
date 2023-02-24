package main
import (
	"flag"
	"log"
	"net"
	"fmt"
	"net/http"
)

var (
	url string
	conn *grpc.ClientConn
)


func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/listar", func(w http.ResponseWriter, r *http.Request) {
		response, err := m1.list(conn)
		if err != nil {
			log.Fatalf("Error when calling func: %s", err)
		}
		log.Printf("Response from server: %s", response.Body)
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/buscar", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	return mux
}

func runHttp(listenAddr string) error {
	s := http.Server{
		Addr:    listenAddr,
		Handler: router.NewRouter(url), // Our own instance of servemux
	}
	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)
	return s.ListenAndServe()
}

func main() {
	var (
		host = flag.String("host", "", "host http address to listen on")
		port = flag.String("port", "8000", "port number for http listener")
	)
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	
	conn, err := m1.createConnection("localhost:9000")

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	if err := runHttp(addr, url); err != nil {
		log.Fatal(err)
	}

}