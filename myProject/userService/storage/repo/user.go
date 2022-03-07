package repo

import (
	pb "myProject/userService/genproto/user_service"
)

type UserStorageI interface {
	Create(*pb.UserRes) (*pb.UserReq, error)
	GetUser(*pb.ById) (*pb.UserReq, error)
	Update(*pb.UserReq) error
	Delete(*pb.ById) (*pb.Mess, error)
	ListUsers(*pb.Mess) (*pb.ListUser, error)
	CheckField(*pb.Checkfild) (*pb.Mess, error)
	GetByEmail(*pb.GetByemail) (*pb.UserReq, error)
	SqlBld(*pb.SqlbldRes) (*pb.Mess, error)
}
