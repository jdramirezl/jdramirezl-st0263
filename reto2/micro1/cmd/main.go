package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	m1 "retos/reto2/micro1/internal"

	"google.golang.org/grpc"
)

var (
	dir string
)

type FileServiceServer struct {
	m1.UnimplementedFileServiceServer
}

func (file *FileServiceServer) List(ctx context.Context, req *m1.ListRequest) (*m1.FileResponse, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileNames = append(fileNames, file.Name())
	}

	res := &m1.FileResponse{Name: fileNames}
	return res, nil
}

func (file *FileServiceServer) Search(ctx context.Context, req *m1.SearchRequest) (*m1.FileResponse, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var found bool = false
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name() == req.Name {
			found = true
			break
		}
	}

	var res *m1.FileResponse
	if found {
		res = &m1.FileResponse{Name: []string{"Found"}}
	} else {
		res = &m1.FileResponse{Name: []string{}}
	}

	return res, nil
}

func start() *Configuration {
	fmt.Println("Loading configuration")
	config, err := loadConfig("../config")
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return nil
	}
	fmt.Println("Configuration loaded succesfully")
	return config
}

func main() {
	fmt.Println("Starting gRPC server")
	config := start()
	dir = config.Directory

	addr := net.JoinHostPort(config.IP, config.Port)

	fmt.Printf("Starting listener at %s\n", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Listener started succesfully")

	fileServer := &FileServiceServer{}
	grpcServer := grpc.NewServer()

	m1.RegisterFileServiceServer(grpcServer, fileServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	fmt.Printf("Server started succesfully")
}
