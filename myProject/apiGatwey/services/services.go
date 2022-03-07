package services

import (
	"myProject/apiGatwey/config"
	pbt "myProject/apiGatwey/genproto/task_service"
	pbu "myProject/apiGatwey/genproto/user_service"
	pbe "myProject/apiGatwey/genproto/email_service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/credentials/insecure"
)


type IServiceManager interface {
	TaskService() pbt.TaskServiceClient
	UserService() pbu.UserServiceClient
	EmailService() pbe.EmailServiceClient
}

type serviceManager struct {
	taskService pbt.TaskServiceClient
	userService pbu.UserServiceClient
	emailService pbe.EmailServiceClient
}

func (s *serviceManager) TaskService() pbt.TaskServiceClient {
	return s.taskService
}
func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}
func (s *serviceManager) EmailService() pbe.EmailServiceClient {
	return s.emailService
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
	connEmail, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.EmailServiceHost, conf.EmailServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		taskService: pbt.NewTaskServiceClient(connTask),
		userService: pbu.NewUserServiceClient(connUser),
		emailService: pbe.NewEmailServiceClient(connEmail),
	}

	return serviceManager, nil
}
