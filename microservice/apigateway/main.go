package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	port = "8080"
)

type frontendServer struct {
	listSvcAddr string
	listSvcConn *grpc.ClientConn

	addSvcAddr string
	addSvcConn *grpc.ClientConn

	updateSvcAddr string
	updateSvcConn *grpc.ClientConn

	deleteSvcAddr string
	deleteSvcConn *grpc.ClientConn
}

func main() {

	ctx := context.Background()
	svc := new(frontendServer)

	mustMapEnv(&svc.listSvcAddr, "LIST_SERVICE_ADDR")

	mustConnGRPC(ctx, &svc.listSvcConn, svc.listSvcAddr)

	messages, _ := svc.listMessages(ctx)

	fmt.Println("********************")
	fmt.Println(messages)

}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure())
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}
