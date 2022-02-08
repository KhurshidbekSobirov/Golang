package repo

import (
	pb "myProject/userService/genproto/user_service"
)

type UserStorageI interface {
	Create(*pb.UserRes) (*pb.UserReq, error)
	GetUser(*pb.ById) (*pb.UserReq, error)
	Update(*pb.UserReq) (*pb.Mess, error)
	Delete(*pb.ById) (*pb.Mess, error)
}
