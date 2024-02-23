package notificationClient

import (
	"context"
	"fmt"
	"notification-service/errHandler"
	"notification-service/pb"
)

func CallNotificationExists(client pb.NotificationServiceClient) {
	request := &pb.NotificationExistsRequest{
		Pid: "",
	}

	response, err := client.NotificationExists(context.Background(), request)
	errHandler.ErrHandler(err, "Error calling notification exists")

	fmt.Println(response)
}

func CallCreateNotification(client pb.NotificationServiceClient) {
	notification := &pb.Notification{}

	request := &pb.NotificationRequest{
		Notification: notification,
	}

	response, err := client.CreateNotification(context.Background(), request)
	errHandler.ErrHandler(err, "Error calling notification create")

	fmt.Println(response)
}

func CallGetNotificationByService(client pb.NotificationServiceClient) {
	request := &pb.NotificationByServiceRequest{
		Service: "",
	}

	response, err := client.GetNotificationByService(context.Background(), request)
	errHandler.ErrHandler(err, "Error calling get notification by service")

	fmt.Println(response)
}

func CallGetAllNotifications(client pb.NotificationServiceClient) {
	request := &pb.NoParam{}

	response, err := client.GetAllNotifications(context.Background(), request)
	errHandler.ErrHandler(err, "Error calling get all notifications")

	fmt.Println(response)
}

func CallGetNotification(client pb.NotificationServiceClient) {
	request := &pb.NotificationCUDRequest{
		Pid: "",
	}

	response, err := client.GetNotification(context.Background(), request)
	errHandler.ErrHandler(err, "Error getting notification")

	fmt.Println(response)
}

func CallDeleteNotification(client pb.NotificationServiceClient) {
	request := &pb.NotificationCUDRequest{
		Pid: "",
	}

	response, err := client.GetNotification(context.Background(), request)
	errHandler.ErrHandler(err, "Error delete notification")

	fmt.Println(response)
}
