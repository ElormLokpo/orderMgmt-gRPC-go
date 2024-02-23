package main

import (
	notificationClient "notification-service/client/notification-client"
	"notification-service/errHandler"

	"notification-service/pb"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	errHandler.ErrHandler(err, "Error Dailing Grpc")

	client := pb.NewNotificationServiceClient(connection)

	notificationClient.CallNotificationExists(client)
}
