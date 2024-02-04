package main

import (
	ProductClient "product-service/client/product-client"
	"product-service/errhandler"
	"product-service/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	errhandler.ErrHandler(err)

	client := pb.NewProductServiceClient(connection)

	ProductClient.CallCreateProduct(client)
}
