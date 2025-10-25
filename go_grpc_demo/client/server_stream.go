package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "com.shreyash/grpc_demo/v2/proto"
)
func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Streaming started with names: %v", names.Names)
	ctx,cancel := context.WithTimeout(context.Background(),time.Minute)
	defer cancel()

	stream,err := client.SayHelloServerSideStreaming(ctx,names)
	if err != nil{
		log.Fatalf("Error while calling SayHelloServerSideStreaming: %v", err)
	}

	for{
		message,err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("Error while streaming: %v",err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
	
	

}