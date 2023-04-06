package main

import (
	"log"
	"net"

	"github.com/golang/protobuf/protoc-gen-go/grpc"
)

type server struct {
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")

	if err != nil {
		log.Fatal("Err while create listen %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	err = s.Server(lis)

	if err != nil {
		log.Fatal("Err while Server $v", err)
	}

}
