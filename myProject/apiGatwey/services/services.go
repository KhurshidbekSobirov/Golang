package services

import (
	"myProject/apiGatwey/config"
	pbt "myProject/apiGatwey/genproto/task_service"
	pbu "myProject/apiGatwey/genproto/user_service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/credentials/insecure"
)


type IServiceManager interface {
	TaskService() pbt.TaskServiceClient
	UserService() pbu.UserServiceClient
}

type serviceManager struct {
	taskService pbt.TaskServiceClient
	userService pbu.UserServiceClient
}

func (s *serviceManager) TaskService() pbt.TaskServiceClient {
	return s.taskService
}
func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connTask, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.TaskServiceHost, conf.TaskServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		taskService: pbt.NewTaskServiceClient(connTask),
		userService: pbu.NewUserServiceClient(connUser),
	}

	return serviceManager, nil
}
