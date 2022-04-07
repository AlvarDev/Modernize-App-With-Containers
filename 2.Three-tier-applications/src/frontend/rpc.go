package main

import (
	"context"
	pb "frontendservice/pb"
)

func (fe *frontendServer) addRemainder(ctx context.Context, req *pb.AddRemainderRequest) (*pb.AddRemainderResponse, error) {
	return pb.NewReminderServiceClient(fe.backendSvcConn).AddRemainder(ctx, req)
}

func (fe *frontendServer) deleteRemainder(ctx context.Context, req *pb.DeleteRemainderRequest) (*pb.DeleteRemainderResponse, error) {
	return pb.NewReminderServiceClient(fe.backendSvcConn).DeleteRemainder(ctx, req)
}

func (fe *frontendServer) listRemainders(ctx context.Context, req *pb.ListRemaindersRequest) (*pb.ListRemaindersResponse, error) {
	return pb.NewReminderServiceClient(fe.backendSvcConn).ListRemainders(ctx, req)
}
