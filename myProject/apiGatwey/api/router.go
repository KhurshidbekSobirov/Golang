package api

import (
	
	v1 "myProject/apiGatwey/api/handler"
	"myProject/apiGatwey/config"
	"myProject/apiGatwey/pkg/logger"
	"myProject/apiGatwey/services"
	docs "myProject/apiGatwey/api/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// @BasePath /v1
// New ...

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/v1"

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/tasks", handlerV1.CreateTask)
	api.POST("/users", handlerV1.CreateUser)
	api.GET("/task/:id", handlerV1.GetTask)
	api.GET("/user/:id",handlerV1.GetUser)
	api.PUT("/user/:id",handlerV1.UpdateUser)
	api.PUT("/task/:id",handlerV1.UpdateTask)
	api.DELETE("/user/:id",handlerV1.DeleteUser)
	api.DELETE("/task/:id",handlerV1.DeleteTask)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}


