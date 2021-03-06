package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "golang_grpc_gin_jaeger_Y/hello"

	"google.golang.org/grpc"
)

func Run() {

	addr := "127.0.0.1:9999"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Can not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Message: "Moto"})
	if err != nil {
		log.Fatalf("Could not get nonce: %v", err)
	}

	fmt.Println("Response:", r.GetMessage())
}
