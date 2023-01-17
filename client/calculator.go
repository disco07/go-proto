package main

import (
	"context"
	pb "github.com/disco07/calculator/pb"
	"log"
)

func doCalculatorAdd(c pb.CalculatorServiceClient) {
	res, err := c.Add(context.Background(), &pb.CalculatorRequest{
		NumberOne: 10.5,
		NumberTwo: 12,
	})
	if err != nil {
		log.Fatalf("could not sum :%v", err)
	}
	log.Printf("Sum %v", res.Result)
}
