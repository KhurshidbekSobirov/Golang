package handlers

import (
	"errors"
	"myProject/apiGatwey/api/token"
	"myProject/apiGatwey/api/models"
	"myProject/apiGatwey/config"
	"myProject/apiGatwey/pkg/logger"
	"myProject/apiGatwey/services"
	"myProject/apiGatwey/storage/repo"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log             logger.Logger
	serviceManager  services.IServiceManager
	inMemoryStorage repo.InMemoryStorageI
	cfg             config.Config
	jwtHandler      token.JWTHandler
}

type HandlerV1Config struct {
	Logger          logger.Logger
	ServiceManager  services.IServiceManager
	InMemoryStorage repo.InMemoryStorageI
	Cfg             config.Config
	JwtHandler      token.JWTHandler
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:             c.Logger,
		serviceManager:  c.ServiceManager,
		cfg:             c.Cfg,
		inMemoryStorage: c.InMemoryStorage,
	}
}

func GetClaims(h *handlerV1, c *gin.Context) jwt.MapClaims {
	var (
		ErrUnauthorized = errors.New("unauthorized")
		authorization   models.GetProfileByJwtRequestModel
		claims          jwt.MapClaims
		err             error
	)
	authorization.Token = c.GetHeader("Authorization")
	strSlice := strings.Split(authorization.Token, " ")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Unauthorized request",
			},
		})
		h.log.Error("Unauthorized request", logger.Error(ErrUnauthorized))
		return nil
	}
	h.jwtHandler.Token = strSlice[1]
	claims, err = h.jwtHandler.ExtractClaims()
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Unauthorized request",
			},
		})
	}
	return claims
}
