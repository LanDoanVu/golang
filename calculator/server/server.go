package main

import (
	"fmt"
	"log"
	"net"

	"github.com/LanDoanVu/golang/golang/calculator/calculatorpb"
	"google.golang.org/grpc"
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

	fmt.Printf("Calculator is running......")

	err = s.Serve(lis)

	if err != nil {
		log.Fatal("Err while Server $v", err)
	}

}
