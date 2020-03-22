package service

import (
	"context"
	pb "github.com/hitsuji-haneta/playk8s/sub/protocol"
)

type GreetService struct {
}

func (s *GreetService) Hello(ctx context.Context, req *pb.Greeter) (*pb.GreetingResponse, error) {
	return &pb.GreetingResponse{Greeting: "hello", Name: req.Name}, nil
}
