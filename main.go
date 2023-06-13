package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ahfrd/grpc/micro-topup/config"
	"github.com/ahfrd/grpc/micro-topup/src/proto/topup"
	"github.com/ahfrd/grpc/micro-topup/src/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	fmt.Println(c.Port)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Emoney Svc on", c.Port)

	s := services.TopUpService{}

	grpcServer := grpc.NewServer()

	topup.RegisterTopUpServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
