package main

import (
	"auth-service/errHandler"
	"auth-service/pb"
	"net"

	authserver "auth-service/server/auth-server"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	errHandler.StdErr(err)

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &authserver.AuthServerStruct{})

	if err = grpcServer.Serve(listener); err != nil {
		errHandler.StdErr(err)
	}
}
