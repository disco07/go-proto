package main

import (
	"context"
	pb "github.com/disco07/calculator/pb"
)

func (s *Server) Add(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {

	return &pb.CalculatorResponse{
		Result: in.GetNumberTwo() + in.GetNumberOne(),
	}, nil
}
