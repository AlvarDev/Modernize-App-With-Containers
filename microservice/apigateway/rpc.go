package main

import (
	"context"
	pb "mymessagesapigateway/pb"
)

func (fe *frontendServer) listMessages(ctx context.Context) (*pb.ListMyMessagesResponse, error) {
	return pb.NewMyMessageServiceClient(fe.listSvcConn).
		ListMessages(ctx, &pb.ListMyMessagesRequest{})
}
