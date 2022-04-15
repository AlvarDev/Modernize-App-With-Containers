package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "listservice/pb"

	"google.golang.org/grpc"
)

const (
	port = ":50054"
)

type server struct {
	pb.UnimplementedReminderServiceServer
}

func (s *server) ListRemainders(ctx context.Context, req *pb.ListRemaindersRequest) (*pb.ListRemaindersResponse, error) {
	remainders, err := ListRemainders(req.GetUserUID())
	return &pb.ListRemaindersResponse{Remainders: remainders}, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listening: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterReminderServiceServer(s, &server{})
	fmt.Printf("server listening at: %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}
