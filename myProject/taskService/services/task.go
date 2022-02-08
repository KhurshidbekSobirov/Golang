package services

import (
	"context"

	pb "myProject/taskService/genproto/task_service"
	l "myProject/taskService/pkg/logger"
	"myProject/taskService/storage"

	"github.com/jmoiron/sqlx"
)

type taskService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewTaskService(db *sqlx.DB, log l.Logger) *taskService {
	return &taskService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *taskService) Create(ctx context.Context, req *pb.TaskRes) (*pb.TaskReq, error) {
	res, err := s.storage.Task().Create(req)
	if err != nil {
		s.logger.Error("error while creating a task",l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *taskService) GetTask(ctx context.Context, req *pb.ById) (*pb.TaskReq, error) {
	res, err := s.storage.Task().GetTask(req)
	if err != nil {
		s.logger.Error("error while getting a task",l.Error(err))
		return nil, err
	}
	return res, nil

}

func (s *taskService) Update(ctx context.Context, req *pb.TaskReq) (*pb.Mess, error) {
	res, err := s.storage.Task().UpdateTask(req)
	if err != nil {
		s.logger.Error("error while update a task",l.Error(err))
		return nil, err
	}
	return res, nil
}

func (s *taskService) Delete(ctx context.Context, req *pb.ById) (*pb.Mess, error) {
	res, err := s.storage.Task().DeleteTask(req)
	if err != nil {
		s.logger.Error("error while delete a task",l.Error(err))
		return nil, err
	}
	return res, nil
}
func (s *taskService) ListOverdue(ctx context.Context, req *pb.Mess) (*pb.ListTasks, error) {
	res, err := s.storage.Task().ListOverdue(nil)
	if err != nil {
		s.logger.Error("error while listing a task",l.Error(err))
		return nil, err
	}
	return res, nil
}


func (r *taskService) GetTasks(ctx context.Context, user *pb.ByUserId) (*pb.ListTasks, error) {

	res, err := r.storage.Task().ListTasks(user)
	if err != nil {
		r.logger.Error("error while listing a task",l.Error(err))
		return nil, err
	}
	return res, nil	
}