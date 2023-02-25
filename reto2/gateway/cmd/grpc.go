package main

import (
	m1 "retos/reto2/gateway/internal"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Search(conn *grpc.ClientConn, _name string) (*m1.FileResponse, error) {
	client := m1.NewFileServiceClient(conn)
	searchreq := m1.SearchRequest{Name: _name}
	return client.Search(context.Background(), &searchreq)
}

func List(conn *grpc.ClientConn) (*m1.FileResponse, error) {
	client := m1.NewFileServiceClient(conn)
	listrequest := m1.ListRequest{}
	res, err := client.List(context.Background(), &listrequest)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateConnection(url string) (*grpc.ClientConn, error) {
	return grpc.Dial(url, grpc.WithInsecure())
}
