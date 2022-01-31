package grpcclient

import "Golang/microservice/storage"

type TaskService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewTaskService(db *sqlx.DB, log l.Logger) *TaskService {
	return &TaskService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *TaskService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	res, err := s.storage.User().Create(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TaskService) ListUsers(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	res, err := s.storage.User().ListUsers(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TaskService) GetUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	res, err := s.storage.User().GetUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *TaskService) DeleteUser(ctx context.Context, req *pb.User) (*pb.Xabar, error) {
	res, err := s.storage.User().DeleteUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TaskService) UpdateUser(ctx context.Context, req *pb.User) (*pb.Xabar, error) {
	res, err := s.storage.User().UpdateUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TaskService) Search(ctx context.Context, req *pb.SearchUser) (*pb.User, error) {
	res, err := s.storage.User().Search(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
