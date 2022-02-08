package services

import (
	"context"
	pbt "myProject/userService/genproto/task_service"
	pb "myProject/userService/genproto/user_service"
	l "myProject/userService/pkg/logger"
	"myProject/userService/services/grpcClient"
	"myProject/userService/storage"

	"github.com/jmoiron/sqlx"
)

type UserService struct {
	storage storage.IStorage
	client  grpcClient.IService
	logger  l.Logger
}

func NewUserService(db *sqlx.DB, client grpcClient.IService, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		client:  client,
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.UserRes) (*pb.UserReq, error) {
	res, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("error while creating a user", l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.ById) (*pb.UserReq, error) {
	res, err := s.storage.User().GetUser(req)
	if err != nil {
		s.logger.Error("error while getting a user", l.Error(err))
		return nil, err
	}

	tasks, err := s.client.TaskService().GetTasks(ctx, &pbt.ByUserId{
		UserId: req.UserId,
	})

	if err != nil {
		s.logger.Error("error while getting tasks", l.Error(err))
		return nil, err
	}

	//var ts []*pb.TaskRes

	for _, val := range tasks.Tasks {
		res.Tasks = append(res.Tasks, &pb.TaskRes{
			Id:         val.Id,
			Name:       val.Name,
			Deadline:   val.Deadline,
			Summary:    val.Summary,
			Status:     val.Status,
			AssigneeId: val.AssigneId,
		})
	}

	return res, nil
}

func (s *UserService) Update(ctx context.Context, req *pb.UserReq) (*pb.Mess, error) {
	res, err := s.storage.User().Update(req)

	if err != nil {
		s.logger.Error("error while updating a user", l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *UserService) Delete(ctx context.Context, req *pb.ById) (*pb.Mess, error) {
	res, err := s.storage.User().Delete(req)

	if err != nil {
		s.logger.Error("error while deleting a user", l.Error(err))
		return nil, err
	}
	return res, nil
}
