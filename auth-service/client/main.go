package main

import (
	authclient "auth-service/client/auth-client"
	"auth-service/errHandler"
	"auth-service/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	errHandler.StdErr(err)

	client := pb.NewAuthServiceClient(connection)
	authclient.CallRegisterUser(client)
	authclient.CallLoginUser(client)

}
