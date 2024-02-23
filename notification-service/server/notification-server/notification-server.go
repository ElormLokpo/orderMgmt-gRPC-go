package notificationServer

import (
	"context"
	"notification-service/errHandler"
	"notification-service/pb"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func (server NotificationServerStruct) NotificationExists(ctx context.Context, req *pb.NotificationExistsRequest) (*pb.NotificationExistsResponse, error) {
	notificationId := req.GetPid()

	filter := bson.M{"pid": notificationId}

	var foundNotification *pb.Notification

	err := server.collection.FindOne(context.Background(), filter).Decode(&foundNotification)
	errHandler.ErrHandler(err, "Error finding notification Id")

	res := &pb.NotificationExistsResponse{
		Exists: true,
	}

	return res, nil

}

func (server NotificationServerStruct) CreateNotification(ctx context.Context, req *pb.NotificationRequest) (*pb.NotificationCUDResponse, error) {
	notification := req.GetNotification()

	notification.Pid = uuid.NewString()

	_, err := server.collection.InsertOne(context.Background(), notification)
	errHandler.ErrHandler(err, "Error adding notification")

	res := &pb.NotificationCUDResponse{
		Message: "Notification created successfully",
	}

	return res, nil
}

func (server NotificationServerStruct) GetAllNotifications(ctx context.Context, req *pb.NoParam) (*pb.NotificationArrResponse, error) {

	data, err := server.collection.Find(context.Background(), bson.M{})
	errHandler.ErrHandler(err, "Error fetching all notifications")

	var notificationsArr []*pb.Notification

	for data.Next(context.Background()) {
		var notification *pb.Notification

		if err := data.Decode(&notification); err != nil {
			errHandler.ErrHandler(err, "Error fetching notification")
		}

		notificationsArr = append(notificationsArr, notification)
	}

	res := &pb.NotificationArrResponse{
		Notifications: notificationsArr,
	}

	return res, nil
}

func (server NotificationServerStruct) GetNotification(ctx context.Context, req *pb.NotificationCUDRequest) (*pb.NotificationResponse, error) {
	notificationId := req.GetPid()

	var notification *pb.Notification

	filter := bson.M{"pid": notificationId}
	err := server.collection.FindOne(context.Background(), filter).Decode(&notification)
	errHandler.ErrHandler(err, "Error finding notification")

	res := &pb.NotificationResponse{
		Notification: notification,
	}

	return res, nil
}

func (server NotificationServerStruct) GetNotificationByService(ctx context.Context, req *pb.NotificationByServiceRequest) (*pb.NotificationResponse, error) {
	notificationService := req.GetService()

	var notification *pb.Notification

	filter := bson.M{"service": notificationService}
	err := server.collection.FindOne(context.Background(), filter).Decode(&notification)
	errHandler.ErrHandler(err, "Error finding notification")

	res := &pb.NotificationResponse{
		Notification: notification,
	}

	return res, nil
}

func (server NotificationServerStruct) DeleteNotification(ctx context.Context, req *pb.NotificationCUDRequest) (*pb.NotificationCUDResponse, error) {
	notificationId := req.GetPid()

	var notification *pb.Notification

	filter := bson.M{"pid": notificationId}
	err := server.collection.FindOneAndDelete(context.Background(), filter).Decode(&notification)
	errHandler.ErrHandler(err, "Error finding notification")

	res := &pb.NotificationCUDResponse{
		Message:      "Notification deleted successfully",
		Notification: notification,
	}

	return res, nil
}
