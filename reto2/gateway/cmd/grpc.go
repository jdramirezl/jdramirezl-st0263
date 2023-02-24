package m1

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func search()(m1.GetFilesResponse, error){
	client := m1.NewFileServiceClient(conn)
	listrequest := m1.ListRequest{}
	return client.List(context.Background(), listrequest)
}

func list()(m1.GetFilesResponse, error){
	client := m1.NewFileServiceClient(conn)
	listrequest := m1.ListRequest{}
	return client.List(context.Background(), listrequest)
}

func createConnection(url string) (*grpc.ClientConn, error){
	return grpc.Dial(url, grpc.WithInsecure())
}