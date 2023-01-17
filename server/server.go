package main

import (
	"fmt"
	pb "github.com/disco07/calculator/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const grpcPort = "50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", grpcPort, err)
	}
	log.Printf("listening on %s", grpcPort)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("error serving %s: %v", grpcPort, err)
	}
}
