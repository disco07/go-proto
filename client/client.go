package main

import (
	"fmt"
	pb "github.com/disco07/calculator/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const grpcPort = "50051"

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%s", grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	doCalculatorAdd(c)
}
