package main

import (
	"net"
	"product-service/errhandler"
	"product-service/pb"
	ProductServer "product-service/server/product-server"

	"product-service/database"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	errhandler.ErrHandler(err)

	database.ConnectDB()
	grpcServer := grpc.NewServer()

	client := database.GetClient()

	pb.RegisterProductServiceServer(grpcServer, ProductServer.NewProductServiceServer(client))

	if err := grpcServer.Serve(listener); err != nil {
		errhandler.ErrHandler(err)
	}

}
