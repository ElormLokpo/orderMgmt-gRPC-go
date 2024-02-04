package orderclient

import (
	"context"
	"fmt"
	"order-service/errHandler"
	"order-service/pb"
)

func CallPlaceOrder(client pb.OrderServiceClient) {
	req := &pb.Order{}
	res, err := client.PlaceOrder(context.Background(), req)
	errHandler.StdErr(err)

	fmt.Println(res)

}
