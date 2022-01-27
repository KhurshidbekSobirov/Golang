package service

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	pb "github.com/KhurshidbekSobirov/Golang/microservice/genproto"
	l"github.com/KhurshidbekSobirov/Golang/microservice/pkg/logger"
	"github.com/KhurshidbekSobirov/Golang/microservice/storage"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}


func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
		user, err := s.storage.User().Create(req)
		if err != nil {
			fmt.Println(err)
		  s.logger.Error("Error while creating a new  user")
		  return nil, err
		  
		}
		return user, nil
}

