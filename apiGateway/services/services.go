package services

import (
	"Golang/apiGateway/config"
	pb "Golang/apiGateway/genproto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/credentials/insecure"
)

type IServiceManager interface {
	TaskService() pb.TaskServiceClient
}

type serviceManager struct {
	taskService pb.TaskServiceClient
}

func (s *serviceManager) TaskService() pb.TaskServiceClient {
	return s.taskService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connTask, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.TaskServiceHost, conf.TaskServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		taskService: pb.NewTaskServiceClient(connTask),
	}

	return serviceManager, nil
}
