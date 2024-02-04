package main

import (
	"net"
	"order-service/errHandler"
	"order-service/pb"
	oderserver "order-service/server/order-server"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	errHandler.StdErr(err)

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &oderserver.OrderServerStruct{})
	if err = grpcServer.Serve(listener); err != nil {
		errHandler.StdErr(err)
	}
}
