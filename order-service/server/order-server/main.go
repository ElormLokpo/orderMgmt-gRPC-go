package orderserver

import (
	"context"
	"order-service/pb"
)

type OrderServerStruct struct {
	pb.OrderServiceServer
}

func (server *OrderServerStruct) PlaceOrder(context context.Context, req *pb.Order) (*pb.Bill, error) {

	final_bill := &pb.Bill{}

	return final_bill, nil
}
