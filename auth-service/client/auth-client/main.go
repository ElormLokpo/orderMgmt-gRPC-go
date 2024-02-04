package authclient

import (
	"auth-service/errHandler"
	"auth-service/pb"
	"context"
	"fmt"
)

func CallRegisterUser(client pb.AuthServiceClient) {
	sampleUser := &pb.User{
		Pid:      "1",
		Fullname: "Jack Aphese",
		Email:    "jack@gmail.com",
		Password: "12345678",
	}

	req := &pb.RegisterRequest{
		User: sampleUser,
	}

	res, err := client.RegisterUser(context.Background(), req)
	errHandler.StdErr(err)

	fmt.Println(res)

}

func CallLoginUser(client pb.AuthServiceClient) {
	req := &pb.LoginRequest{
		Email:    "john@gmail.com",
		Password: "12345678",
	}

	res, err := client.LoginUser(context.Background(), req)
	errHandler.StdErr(err)

	fmt.Println(res)

}
