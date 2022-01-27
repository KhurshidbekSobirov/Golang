package main

import (
    "github.com/KhurshidbekSobirov/Golang/template-service/config"
    pb "github.com/KhurshidbekSobirov/Golang/template-service/genproto"
    "github.com/KhurshidbekSobirov/Golang/template-service/pkg/db"
    "github.com/KhurshidbekSobirov/Golang/template-service/pkg/logger"
    "github.com/KhurshidbekSobirov/Golang/template-service/service"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "net"
)

func main() {
    cfg := config.Load()

    log := logger.New(cfg.LogLevel, "template-service")
    defer logger.Cleanup(log)

    log.Info("main: sqlxConfig",
        logger.String("host", cfg.PostgresHost),
        logger.Int("port", cfg.PostgresPort),
        logger.String("database", cfg.PostgresDatabase))

    connDB, err := db.ConnectToDB(cfg)
    if err != nil {
        log.Fatal("sqlx connection to postgres error", logger.Error(err))
    }

    userService := service.NewUserService(connDB, log)

    lis, err := net.Listen("tcp", cfg.RPCPort)
    if err != nil {
        log.Fatal("Error while listening: %v", logger.Error(err))
    }

    s := grpc.NewServer()
    reflection.Register(s)
    pb.RegisterUserServiceServer(s, userService)
    log.Info("main: server running",
        logger.String("port", cfg.RPCPort))

    if err := s.Serve(lis); err != nil {
        log.Fatal("Error while listening: %v", logger.Error(err))
    }
}
