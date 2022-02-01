package service

import (
	"context"

	pb "app/genproto"
	l "app/pkg/logger"
	"app/storage"

	"github.com/jmoiron/sqlx"
)

//UserService ...
type taskService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewUserService ...
func NewTaskService(db *sqlx.DB, log l.Logger) *taskService {
	return &taskService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *taskService) Create(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	res, err := s.storage.Task().Create(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *taskService) GetTask(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	res, err := s.storage.Task().GetTask(req)
	if err != nil {
		return nil, err
	}
	return res, nil
	
}

func (s *taskService) Update(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	res, err := s.storage.Task().Update(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *taskService) Delete(ctx context.Context, req *pb.Task) (*pb.Mess, error) {
	res, err := s.storage.Task().Delete(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *taskService) ListOverdue(ctx context.Context, req *pb.Mess) (*pb.ListTask, error) {
	res, err := s.storage.Task().ListOverdue(nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
