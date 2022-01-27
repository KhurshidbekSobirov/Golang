package repo

import (
    pb "github.com/KhurshidbekSobirov/Golang/microservice/genproto"
)

//UserStorageI ...
type UserStorageI interface {
    Create(*pb.User) (*pb.User, error)
}
