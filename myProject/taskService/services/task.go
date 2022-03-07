package services

import (
	"context"
	"fmt"

	pb "myProject/taskService/genproto/task_service"
	pbe "myProject/taskService/genproto/email_service"
	pbt "myProject/taskService/genproto/user_service"

	l "myProject/taskService/pkg/logger"
	"myProject/taskService/services/grpcClient"
	"myProject/taskService/storage"

	"github.com/jmoiron/sqlx"
)

type taskService struct {
	storage storage.IStorage
	client  grpcClient.IService
	logger  l.Logger
}

func NewTaskService(db *sqlx.DB, log l.Logger, client grpcClient.IService) *taskService {
	return &taskService{
		storage: storage.NewStoragePg(db),
		client:  client,
		logger:  log,
	}
}

func (s *taskService) Create(ctx context.Context, req *pb.TaskRes) (*pb.TaskReq, error) {
	res, err := s.storage.Task().Create(req)
	if err != nil {
		s.logger.Error("error while creating a task", l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *taskService) GetTask(ctx context.Context, req *pb.ById) (*pb.TaskReq, error) {
	res, err := s.storage.Task().GetTask(req)
	if err != nil {
		s.logger.Error("error while getting a task", l.Error(err))
		return nil, err
	}

	user, err := s.client.UserService().GetUser(ctx, &pbt.ById{UserId: res.AssigneId})
	
	if err != nil {
		return nil, err
	}
	// res.User.FirstName = &pb.User{FirstName: "salom"}
	fmt.Print("\n\n\n",user.FirstName,"\n\n\n")
	fmt.Print("\n\n\n",res,"\n\n\n")

	res.User.LastName = user.LastName
	res.User.Email = user.Email
	res.User.Username = user.Username
	res.User.Gender = user.Gender
	res.User.ProfilePhoto = user.ProfilePhoto
	res.User.Bio = user.Bio

	return res, nil

}

func (s *taskService) Update(ctx context.Context, req *pb.TaskReq) (*pb.Mess, error) {
	res, err := s.storage.Task().UpdateTask(req)
	if err != nil {
		s.logger.Error("error while update a task", l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *taskService) Delete(ctx context.Context, req *pb.ById) (*pb.Mess, error) {
	res, err := s.storage.Task().DeleteTask(req)
	if err != nil {
		s.logger.Error("error while delete a task", l.Error(err))
		return nil, err
	}
	return res, nil
}
func (s *taskService) ListOverdue(ctx context.Context, req *pb.Mess) (*pb.ListTasks, error) {
	res, err := s.storage.Task().ListOverdue(nil)
	if err != nil {
		s.logger.Error("error while listing a task", l.Error(err))
		return nil, err
	}
	for _, val := range res.Tasks{
		user, err := s.client.UserService().GetUser(ctx, &pbt.ById{UserId: val.AssigneId})
		// if err != nil{
		// 	return nil,err
		// }
		// emails = append(emails ,user.Email)
		

		email := &pbe.Email{
			Id: "",
			Body: val.Name + " task was not completed on time",
			Subject: "Warning!",
			Recipient: user.Email,
		}

		_,err = s.client.EmailService().Send(ctx,email)
		if err != nil{
			return nil, err
		}
		for _,num := range user.PhoneNumbers{
			sms := &pbe.Sms{
				From:     "Sifat Nazorati",
				ApiKey:    "fe89fc68",
				ApiSecret: "wN4fbElg0WWaa7Q5",
				To:       num.Name,
				Text:      val.Name + " task was not completed on time",
			}
			_,errr := s.client.EmailService().SendSms(ctx,sms)
				if errr != nil {
					return nil,err
				}
			}

	}
	return res, nil
}

func (r *taskService) GetTasks(ctx context.Context, user *pb.ByUserId) (*pb.ListTasks, error) {

	res, err := r.storage.Task().ListTasks(user)
	if err != nil {
		r.logger.Error("error while listing a task", l.Error(err))
		return nil, err
	}
	return res, nil
}
