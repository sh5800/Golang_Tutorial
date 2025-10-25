package main

import (
	"log"
	"time"

	pb "com.shreyash/grpc_demo/v2/proto"
)

func (s *helloServer) SayHelloServerSideStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerSideStreamingServer) error {
	log.Printf("Got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{Message: "Hello " + name}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
