package notificationServer

import (
	"notification-service/pb"

	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationServerStruct struct {
	pb.NotificationServiceServer
	collection mongo.Collection
}

func NewNotificationServer(client *mongo.Client) NotificationServerStruct {
	collection := client.Database("orderGrpc").Collection("notifications")

	return NotificationServerStruct{
		collection: *collection,
	}

}
