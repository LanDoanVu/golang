package main

import (
	"context"
	"io"
	"log"

	"github.com/LanDoanVu/golang/grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Erro while conn Dial $v", err)
	}
	defer conn.Close()

	client := calculatorpb.NewCalculatorServiceClient(conn)

	log.Printf("Service client %f", client)

	// callSum(client)

	callSNT(client)

}

func callSum(call calculatorpb.CalculatorServiceClient) {
	log.Println("Calling sum API")
	resp, err := call.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 10,
		Num2: 12,
	})

	if err != nil {
		log.Fatal("Call sum API err %v", err)
	}

	log.Println("Sum API response %v", resp.GetResult())
}

func callSNT(call calculatorpb.CalculatorServiceClient) {

	log.Println("Calling soNT API....")

	stream, err := call.SoNT(context.Background(), &calculatorpb.SoNTRequest{
		Number: 120,
	})

	if err != nil {
		log.Fatalf("Call So Nguyen To err %v", err)
	}

	for {
		resp, recvErr := stream.Recv()
		if recvErr != io.EOF {
			log.Println("Server finish streaming")
			return
		}

		if recvErr != nil {
			log.Fatal("Call soNT recvErr %v", recvErr)

		}

		log.Printf("So Nguyen To %v", resp.GetResult())
	}
}
