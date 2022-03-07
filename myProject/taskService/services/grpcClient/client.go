package grpcClient

import (
	"fmt"
	"myProject/taskService/config"
	pb "myProject/taskService/genproto/user_service"
	pbe "myProject/taskService/genproto/email_service"
	"google.golang.org/grpc"
)

// IService interface for task service client implementing
type IService interface {
	UserService() pb.UserServiceClient
	EmailService() pbe.EmailServiceClient
}

type GrpcClient struct {
	cfg config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (IService, error) {
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserserviceHost, cfg.UserservicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%d, err:%s", 
	cfg.UserserviceHost, cfg.UserservicePort, err)
	}

	connEmail, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.EmailserviceHost, cfg.EmailservicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%d, err:%s", 
	cfg.EmailserviceHost, cfg.EmailservicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"user_service": pb.NewUserServiceClient(connUser),
		
			"email_service": pbe.NewEmailServiceClient(connEmail),
		},
	}, nil
}

func (g *GrpcClient) UserService() pb.UserServiceClient {
	return g.connections["user_service"].(pb.UserServiceClient)
}

func (g *GrpcClient) EmailService() pbe.EmailServiceClient {
	return g.connections["email_service"].(pbe.EmailServiceClient)
}