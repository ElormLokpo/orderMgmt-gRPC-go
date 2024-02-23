package main

import (
	"net"

	"notification-service/database"

	"notification-service/errHandler"
	"notification-service/pb"
	notificationServer "notification-service/server/notification-server"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "8000")
	errHandler.ErrHandler(err, "Error listening on port 8000, notification grpc server")

	database.ConnectDB()
	grpcServer := grpc.NewServer()
	client := database.GetClient()

	pb.RegisterNotificationServiceServer(grpcServer, notificationServer.NewNotificationServer(client))

	if err := grpcServer.Serve(listener); err != nil {
		errHandler.ErrHandler(err, "Err serving grpc server, notification grpc server")
	}
}
