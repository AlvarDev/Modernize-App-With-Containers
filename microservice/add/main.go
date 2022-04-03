package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	pb "mymessagesadd/protos"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
	max  = 1000
	min  = 2
)

// server is used to implement RPC server
type server struct {
	pb.UnimplementedMyMessageServiceServer
}

func (s *server) AddMessage(ctx context.Context, req *pb.AddMyMessageRequest) (*pb.AddMyMessageResponse, error) {
	fmt.Println("Request received!")
	newId := rand.Intn(max-min) + min

	return &pb.AddMyMessageResponse{
		MyMessage: &pb.MyMessage{
			UserId:      req.GetMyMessage().GetUserId(),
			MyMessageId: int64(newId),
			MyMessage:   req.GetMyMessage().GetMyMessage()},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listening: %v:", err)
	}

	s := grpc.NewServer()
	pb.RegisterMyMessageServiceServer(s, &server{})
	fmt.Printf("server listening at: %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
