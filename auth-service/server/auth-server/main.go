package authserver

import (
	"auth-service/pb"
	"auth-service/utilities/bcrypt"
	"auth-service/utilities/jwt"
	"context"
	"fmt"
)

type AuthServerStruct struct {
	pb.AuthServiceServer
}

func (server *AuthServerStruct) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashedPassword := bcrypt.EncryptPassword(req.User.Password)
	fmt.Println(hashedPassword)

	res := &pb.RegisterResponse{
		Userid:  req.User.Pid,
		Message: "Register successful",
	}

	return res, nil
}

func (server *AuthServerStruct) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	sampleUser := &pb.User{
		Pid:      "Sample User 1",
		Fullname: "John Quaye",
		Email:    "john@gmail.com",
		Password: "12345678",
	}

	is_password_valid := bcrypt.CompareHashedPasswords(sampleUser.Password, req.Password)

	fmt.Println(is_password_valid)
	var token string
	var res *pb.LoginResponse

	if is_password_valid == true {
		token = jwt.GenerateToken(sampleUser.Pid)

		res = &pb.LoginResponse{
			Userid:  sampleUser.Pid,
			Token:   token,
			Message: "Login successful",
		}
	} else {
		res = &pb.LoginResponse{
			Userid:  sampleUser.Pid,
			Token:   "No token",
			Message: "Wrong password, Login unsuccessful!",
		}
	}

	return res, nil
}
