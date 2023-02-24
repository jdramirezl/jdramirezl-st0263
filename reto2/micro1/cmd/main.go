package main

import (
	"fmt"
	"log"
	"net"

	"retos/reto2/micro1/m1"
	"google.golang.org/grpc"
)

type FileServiceServer struct{
	m1.UnimplementedFileServiceServer
	directory string
}

func (file *FileServiceServer) List(ctx context.Context, req *m1.ListRequest) (*m1.GetFilesResponse , error) {

	return "xd", nil
}

func (file *FileServiceServer) Search(ctx context.Context, req *m1.SearchRequest) (*m1.GetFileResponse, error) {

	return nil, nil
}

func main() {

	// Update with env
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fileServer := &FileServiceServer{}
	grpcServer := grpc.NewServer()

	m1.RegisterFileServiceServer(grpcServer, fileServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}