package main

import (
	"context"

	pb "com.shreyash/grpc_demo/v2/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello World"}, nil
}
