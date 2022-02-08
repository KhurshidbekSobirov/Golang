package api

import (
	"Golang/apiGateway/config"
	"Golang/apiGateway/pkg/logger"
	"Golang/apiGateway/services"
ginSwagger	"github.com/swaggo/gin-swagger" // gin-swagger middleware
swaggerfiles	 "github.com/swaggo/files"
	v1 "Golang/apiGateway/api/handlers"
docs "Golang/apiGateway/api/docs"
	"github.com/gin-gonic/gin"
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
	docs.SwaggerInfo.BasePath = "/v1"

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})
	

	api := router.Group("/v1")
	api.POST("/users", handlerV1.CreateTask)
	api.GET("/users/:id", handlerV1.GetTask)
	api.PUT("/users", handlerV1.UpdateTask)
	api.DELETE("/users/:id", handlerV1.DeleteTask)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
