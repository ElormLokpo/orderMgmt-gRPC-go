package main

import (
	"net"
	"product-service/errhandler"
	"product-service/pb"
	ProductServer "product-service/server/product-server"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "8000")
	errhandler.ErrHandler(err)

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &ProductServer.ProductServiceServerStruct{})

	if err := grpcServer.Serve(listener); err != nil {
		errhandler.ErrHandler(err)
	}

}
