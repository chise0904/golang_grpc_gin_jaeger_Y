package grpcServer

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang_grpc_gin_jaeger_Y/EchoServer"
	"golang_grpc_gin_jaeger_Y/client"

	pb "golang_grpc_gin_jaeger_Y/hello"

	"google.golang.org/grpc"
)

func Run() {

	apiListener, err := net.Listen("tcp", ":9997")
	if err != nil {
		log.Println(err)
		return
	}

	// 註冊 grpc
	es := &EchoServer.EchoServer{}

	grpc := grpc.NewServer()

	pb.RegisterGreeterServer(grpc, es)

	go runClient()

	if err := grpc.Serve(apiListener); err != nil {
		log.Fatal(" grpc.Serve Error: ", err)
		return
	}

}

func runClient() {

	time.Sleep(10 * time.Second)
	fmt.Println("After 3 seconds")

	client.Run()
}
