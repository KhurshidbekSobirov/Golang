package api

import (
	docs "myProject/apiGatwey/api/docs"
	v1 "myProject/apiGatwey/api/handler"
	"myProject/apiGatwey/config"
	"myProject/apiGatwey/pkg/logger"
	"myProject/apiGatwey/services"
	"myProject/apiGatwey/storage/repo"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf            config.Config
	Logger          logger.Logger
	ServiceManager  services.IServiceManager
	InMemoryStorage repo.InMemoryStorageI
}

// @BasePath /v1
// New ...
// @SecurityDefinitions.apikey BearerAuth
// @Description GetMyProfile
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/v1"

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:          option.Logger,
		ServiceManager:  option.ServiceManager,
		Cfg:             option.Conf,
		InMemoryStorage: option.InMemoryStorage,
	})

	api := router.Group("/v1")
	api.POST("/tasks", handlerV1.CreateTask)
	api.POST("/users", handlerV1.CreateUser)
	api.GET("/task/:id", handlerV1.GetTask)
	api.GET("/user/:id", handlerV1.GetUser)
	api.PUT("/user/:id", handlerV1.UpdateUser)
	api.PUT("/task/:id", handlerV1.UpdateTask)
	api.DELETE("/user/:id", handlerV1.DeleteUser)
	api.DELETE("/task/:id", handlerV1.DeleteTask)
	api.GET("/tasks", handlerV1.ListOverdue)
	api.POST("/register", handlerV1.Register)
	api.GET("/profile", handlerV1.GetMyProfile)
	api.POST("/verify/:code", handlerV1.Verify)
	api.PUT("/login",handlerV1.LogIn)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
