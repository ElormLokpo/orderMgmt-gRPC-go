package server

import (
	"context"
	"fmt"
	"product-service/pb"

	"github.com/google/uuid"
)

type ProductServiceServerStruct struct {
	pb.ProductServiceServer
}

func (server ProductServiceServerStruct) CreateProduct(context context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	productFromRequest := req.GetProduct()

	productFromRequest.Pid = uuid.NewString()

	response := &pb.ProductResponse{
		Product: productFromRequest,
	}

	return response, nil

}

func (server ProductServiceServerStruct) GetAllProducts(context context.Context, req *pb.NoParam) (*pb.ProductArrRespnose, error) {

	var sampleProductsArr []*pb.Product

	response := &pb.ProductArrRespnose{
		Product: sampleProductsArr,
	}

	return response, nil

}

func (server ProductServiceServerStruct) GetProductById(context context.Context, req *pb.ProductIdRequest) (*pb.ProductResponse, error) {
	pid := req.GetPid()

	fmt.Println(pid)
	sampleProduct := pb.Product{}

	response := &pb.ProductResponse{
		Product: &sampleProduct,
	}

	return response, nil

}

func (server ProductServiceServerStruct) UpdateProduct(context context.Context, req *pb.ProductIdRequest) (*pb.ProductResponse, error) {
	pid := req.GetPid()

	fmt.Println(pid)
	sampleProduct := pb.Product{}

	response := &pb.ProductResponse{
		Product: &sampleProduct,
	}

	return response, nil

}

func (server ProductServiceServerStruct) DeleteProduct(context context.Context, req *pb.ProductIdRequest) (*pb.ProductResponse, error) {
	pid := req.GetPid()

	fmt.Println(pid)
	sampleProduct := pb.Product{}

	response := &pb.ProductResponse{
		Product: &sampleProduct,
	}

	return response, nil

}
