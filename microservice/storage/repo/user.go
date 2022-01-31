package repo

import (
	pb "github.com/KhurshidbekSobirov/Golang/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	ListUsers(*pb.ListUserRequest) (*pb.ListUserResponse, error)
	GetUser(*pb.User) (*pb.User, error)
	DeleteUser(*pb.User) (*pb.Xabar, error)
	UpdateUser(*pb.User) (*pb.Xabar, error)
	Search(*pb.SearchUser) (*pb.User, error)
}
