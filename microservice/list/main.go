package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "mymessageslist/protos"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement RPC server
type server struct {
	pb.UnimplementedMyMessageServiceServer
}

func (s *server) ListMessages(ctx context.Context, req *pb.ListMyMessagesRequest) (*pb.ListMyMessagesResponse, error) {
	fmt.Println("Request received!!!!")
	return &pb.ListMyMessagesResponse{}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listening: %v:", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyMessageServiceServer(s, &server{})
	fmt.Println("server listening at: %v: ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: %v:", err)
	}

}
