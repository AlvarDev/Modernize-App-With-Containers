package main

import (
	pb "backendservice/pb"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type server struct {
	pb.UnimplementedReminderServiceServer
}

func (s *server) AddRemainder(ctx context.Context, req *pb.AddRemainderRequest) (*pb.AddRemainderResponse, error) {
	newRemainder, err := AddRemainder(req.GetRemainder())
	return &pb.AddRemainderResponse{Remainder: newRemainder}, err
}

func (s *server) DeleteRemainder(ctx context.Context, req *pb.DeleteRemainderRequest) (*pb.DeleteRemainderResponse, error) {
	err := DeleteRemainder(req.GetRemainder())
	return &pb.DeleteRemainderResponse{RemainderID: req.GetRemainder().GetRemainderID()}, err
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
