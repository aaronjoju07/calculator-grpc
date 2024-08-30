package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/aaronjoju07/grpc/pb"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Sum(ctx context.Context, in *pb.NumberRequest) (*pb.CalculationResponse, error) {
	var sum int64
	for _, num := range in.Number {
		sum += num
	}
	return &pb.CalculationResponse{
		Result: sum,
	}, nil
}

func (s *server) Add(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	// Perform the addition
	result := in.GetA() + in.GetB()

	// Return the response
	return &pb.CalculationResponse{
		Result: result,
	}, nil
}

func (s *server) Divide(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	if in.B == 0 {
		return nil, status.Error(
			codes.InvalidArgument, "cannot divide by zero",
		)
	}

	result := in.A / in.B
	return &pb.CalculationResponse{
		Result: result,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to create listner:", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to serve:", err)
	}
}
