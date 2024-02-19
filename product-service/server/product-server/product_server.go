package server

import (
	"context"
	"fmt"
	"product-service/errhandler"
	"product-service/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/google/uuid"
)

type ProductServiceServerStruct struct {
	pb.ProductServiceServer
	collection mongo.Collection
}

func NewProductServiceServer(client *mongo.Client) ProductServiceServerStruct {
	collection := client.Database("productgrpc").Collection("products")

	return ProductServiceServerStruct{
		collection: *collection,
	}
}

func (server ProductServiceServerStruct) CreateProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	productFromRequest := req.GetProduct()

	productFromRequest.Pid = uuid.NewString()

	_, err := server.collection.InsertOne(context.Background(), productFromRequest)
	errhandler.ErrHandler(err)

	response := &pb.ProductResponse{
		Product: productFromRequest,
	}

	return response, nil

}

func (server ProductServiceServerStruct) GetAllProducts(ctx context.Context, req *pb.NoParam) (*pb.ProductArrRespnose, error) {

	var sampleProductsArr []*pb.Product

	data, err := server.collection.Find(context.Background(), bson.M{})
	errhandler.ErrHandler(err)

	for data.Next(context.Background()) {
		var product *pb.Product

		if err := data.Decode(&product); err != nil {
			errhandler.ErrHandler(err)
		}

		sampleProductsArr = append(sampleProductsArr, product)
	}

	response := &pb.ProductArrRespnose{
		Product: sampleProductsArr,
	}

	return response, nil

}

func (server ProductServiceServerStruct) GetProductById(ctx context.Context, req *pb.ProductIdRequest) (*pb.ProductResponse, error) {
	pid := req.GetPid()

	fmt.Println(pid)
	var sampleProduct *pb.Product

	filter := bson.M{"pid": pid}

	err := server.collection.FindOne(context.Background(), filter).Decode(&sampleProduct)
	errhandler.ErrHandler(err)

	response := &pb.ProductResponse{
		Product: sampleProduct,
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

func (server ProductServiceServerStruct) DeleteProduct(ctx context.Context, req *pb.ProductIdRequest) (*pb.DeleteProductResponse, error) {
	pid := req.GetPid()

	filter := bson.M{"pid": pid}
	_, err := server.collection.DeleteOne(context.Background(), filter)
	errhandler.ErrHandler(err)

	response := &pb.DeleteProductResponse{
		Message: fmt.Sprintf("Product %v deleted successfully", pid),
	}

	return response, nil

}
