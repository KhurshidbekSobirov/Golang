package main

import (
	"fmt"
	"myProject/apiGatwey/api"
	"myProject/apiGatwey/config"
	"myProject/apiGatwey/pkg/logger"
	"myProject/apiGatwey/services"
	rds "myProject/apiGatwey/storage/redis"
	"myProject/apiGatwey/storage/repo"

	"github.com/gomodule/redigo/redis"
)

func main() {
	var inMemStrg repo.InMemoryStorageI
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}
	pool := redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort,))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	inMemStrg = rds.NewRedisRepo(&pool)

	server := api.New(api.Option{
		Conf:            cfg,
		Logger:          log,
		ServiceManager:  serviceManager,
		InMemoryStorage: inMemStrg,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
