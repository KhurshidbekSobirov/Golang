package main

import (
	"Golang/microservice/config"
	"Golang/microservice/pkg/db"
	"Golang/microservice/pkg/logger"
	"Golang/microservice/service"
	"Golang/taskservise/pkg/logger"
	"net"
)


func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "template-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host",cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	coonDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	userService := service.NewUserService(coonDB,log)

	lis, err := net.Listen("tcp",cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterUserServiceServer(s. userService)
	log.Info("main: server running",
		logger.String("port:",cfg.RPCPort))

	if err := s.Serve(lis); err != nil{
		log.Fatal("Error while listening: %v", logger.Error(err))
	}


}