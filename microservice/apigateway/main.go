package main

import (
	"context"
	"fmt"
	"time"

	pb "mymessagesapigateway/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	domain := "localhost"
	port := "50051"
	host := domain + ":" + port

	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// return nil, fmt.Errorf("did not connect: %v", err)
		fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create client
	client := pb.NewMyMessageServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	r, err := client.ListMessages(ctx, &pb.ListMyMessagesRequest{})
	if err != nil {
		fmt.Println("Da faq")
		fmt.Println(err)
	}

	fmt.Println(r.GetMyMessages())

}
