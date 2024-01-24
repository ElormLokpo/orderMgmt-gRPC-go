package server

import (
	"fmt"
	"product-service/pb"

	"github.com/google/uuid"
)

type ProductServiceServerStruct struct {
	pb.ProductServiceServer
}

func (server *ProductServiceServerStruct) CreateProduct(req *pb.ProductRequest) (*pb.ProductResponse, error) {
	productFromRequest := req.GetProduct()

	productFromRequest.Pid = uuid.NewString()

	response := &pb.ProductResponse{
		Product: productFromRequest,
	}

	return response, nil

}

func (server *ProductServiceServerStruct) GetAllProducts(req *pb.NoParam) (*pb.ProductArrRespnose, error) {

	var sampleProductsArr []*pb.Product

	response := &pb.ProductArrRespnose{
		Product: sampleProductsArr,
	}

	return response, nil

}

func (server *ProductServiceServerStruct) GetProductById(req *pb.ProductIdRequest) (*pb.ProductResponse, error) {
	pid := req.GetPid()

	fmt.Println(pid)
	sampleProduct := pb.Product{}

	response := &pb.ProductResponse{
		Product: &sampleProduct,
	}

	return response, nil

}

func (server *ProductServiceServerStruct) UpdateProduct(req *pb.ProductIdRequest) (*pb.ProductResponse, error) {
	pid := req.GetPid()

	fmt.Println(pid)
	sampleProduct := pb.Product{}

	response := &pb.ProductResponse{
		Product: &sampleProduct,
	}

	return response, nil

}

func (server *ProductServiceServerStruct) DeleteProduct(req *pb.ProductIdRequest) (*pb.ProductResponse, error) {
	pid := req.GetPid()

	fmt.Println(pid)
	sampleProduct := pb.Product{}

	response := &pb.ProductResponse{
		Product: &sampleProduct,
	}

	return response, nil

}
