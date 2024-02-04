package main

import (
	orderclient "order-service/client/order-client"
	"order-service/errHandler"
	"order-service/pb"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	errHandler.StdErr(err)

	client := pb.NewOrderServiceClient(connection)

	orderclient.CallPlaceOrder(client)
}
