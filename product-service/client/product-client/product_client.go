package client

import (
	"context"
	"fmt"
	"product-service/errhandler"
	"product-service/pb"
)

func CallCreateProduct(client pb.ProductServiceClient) {
	sampleProduct := &pb.Product{}

	req := &pb.ProductRequest{
		Product: sampleProduct,
	}

	res, err := client.CreateProduct(context.Background(), req)
	errhandler.ErrHandler(err)

	fmt.Println(res)
}

func CallGetAllProducts(client pb.ProductServiceClient) {

	req := &pb.NoParam{}

	res, err := client.GetAllProducts(context.Background(), req)
	errhandler.ErrHandler(err)

	fmt.Println(res)
}

func CallUpdateProduct(client pb.ProductServiceClient) {
	req := &pb.ProductIdRequest{}
	res, err := client.UpdateProduct(context.Background(), req)
	errhandler.ErrHandler(err)

	fmt.Println(res)
}

func CallDeleteProduct(client pb.ProductServiceClient) {
	req := &pb.ProductIdRequest{}
	res, err := client.DeleteProduct(context.Background(), req)
	errhandler.ErrHandler(err)

	fmt.Println(res)
}
