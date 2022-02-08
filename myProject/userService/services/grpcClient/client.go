package grpcClient

import (
	"fmt"
	"myProject/userService/config"
	pbt "myProject/userService/genproto/task_service"

	"google.golang.org/grpc"
)

// IService interface for task service client implementing
type IService interface {
	TaskService() pbt.TaskServiceClient
}

type GrpcClient struct {
	cfg config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (IService, error) {
	connTask, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.TaskServiceHost, cfg.TaskServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%d, err:%s", 
	cfg.TaskServiceHost, cfg.TaskServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"task_service": pbt.NewTaskServiceClient(connTask),
		},
	}, nil
}

func (g *GrpcClient) TaskService() pbt.TaskServiceClient {
	return g.connections["task_service"].(pbt.TaskServiceClient)
}