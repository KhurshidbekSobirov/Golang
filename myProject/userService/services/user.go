package services

import (
	"context"
	"fmt"
	pbt "myProject/userService/genproto/task_service"
	pb "myProject/userService/genproto/user_service"
	l "myProject/userService/pkg/logger"
	"myProject/userService/services/grpcClient"
	"myProject/userService/storage"
	"strings"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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
	fmt.Println(req.Id, "\n\n\n\n\n\n\nq")
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
	err := s.storage.User().Update(req)

	if err != nil {
		s.logger.Error("error while updating a user", l.Error(err))
		return &pb.Mess{Message: "Not Update"}, err
	}

	return &pb.Mess{Message: "Ok"}, nil
}

func (s *UserService) Delete(ctx context.Context, req *pb.ById) (*pb.Mess, error) {
	res, err := s.storage.User().Delete(req)

	if err != nil {
		s.logger.Error("error while deleting a user", l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.Mess) (*pb.ListUser, error) {
	res, err := s.storage.User().ListUsers(req)

	if err != nil {
		s.logger.Error("error while deleting a user", l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *UserService) CheckField(ctx context.Context, req *pb.Checkfild) (*pb.Mess, error) {
	mes, err := s.storage.User().CheckField(req)

	if err != nil {
		s.logger.Error("error while updating a user", l.Error(err))
		return &pb.Mess{Message: "Not Update"}, err
	}

	return mes, nil
}

func (s *UserService) GetByEmail(ctx context.Context, req *pb.GetByemail) (*pb.UserReq, error) {
	req.Email = strings.ToLower(req.Email)
	req.Email = strings.TrimSpace(req.Email)
	fmt.Println(req)

	user, err := s.storage.User().GetByEmail(req)
	if err != nil {
		s.logger.Error("error while getting username for login", l.Error(err))
	}
	fmt.Print(user.Password, req.Password, "\n\\n\n\n")
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		s.logger.Error("failed to compare password", l.Error(err))
		return nil, err
	}
	return user, nil
}

func (s *UserService) SqlBld(ctx context.Context, req *pb.SqlbldRes) (*pb.Mess, error) {
	mess, err := s.storage.User().SqlBld(nil)
	fmt.Println(err)
	return mess, err

}
