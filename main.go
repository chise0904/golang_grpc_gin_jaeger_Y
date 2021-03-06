package main

import (
	"golang_grpc_gin_jaeger_Y/grpcServer"
	"golang_grpc_gin_jaeger_Y/httpServer"
)

func main() {

	ch := make(chan struct{})

	go grpcServer.Run()
	go httpServer.Run()

	<-ch
}
