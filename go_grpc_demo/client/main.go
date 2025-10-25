package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "com.shreyash/grpc_demo/v2/proto"
)

const (
	port = ":8080"
)

type helloClient struct{
	pb.GreetServiceClient
}

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Shreyash","Alice","Bob"},
	}

	//callSayHello(client)
	//callSayHelloServerStream(client, names)
	//callSayHelloClientStream(client,names)
	callSayHelloBidirectionalStreaming(client,names)
}
