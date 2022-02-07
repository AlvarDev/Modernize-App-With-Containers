package main

import (
	"context"
	pb "mymessagesapigateway/pb"
)

func (fe *frontendServer) listMessages(ctx context.Context) ([]*pb.MyMessage, error) {
	r, err := pb.NewMyMessageServiceClient(fe.listSvcConn).
		ListMessages(ctx, &pb.ListMyMessagesRequest{})
	return r.GetMyMessages(), err
}
