package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/LanDoanVu/golang/grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumReponse, error) {
	log.Println("Sum called....")
	resp := &calculatorpb.SumReponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}

func (*server) SoNT(req *calculatorpb.SoNTRequest,
	stream calculatorpb.CalculatorService_SoNTServer) error {
	log.Println("So NT called....")
	k := int32(2)
	N := req.GetNumber()

	for N > 1 {
		if N%k == 0 {
			N = N / k
			//send message to Client
			stream.Send(&calculatorpb.SoNTReqonse{
				Result: k,
			})
			log.Println("So nguyen to %v", k)

		} else {
			k++
			log.Println("So k tang len, hien tai la %v", k)
		}
	}
	return nil

}

func main() {
	lis, err := net.Listen("tcp", ":8080")

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
