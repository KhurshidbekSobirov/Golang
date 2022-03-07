package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"myProject/apiGatwey/api/models"
	"myProject/apiGatwey/api/token"
	"myProject/apiGatwey/etc"

	pbe "myProject/apiGatwey/genproto/email_service"
	pb "myProject/apiGatwey/genproto/user_service"
	l "myProject/apiGatwey/pkg/logger"
	"myProject/apiGatwey/storage/redis"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Register ...
// @Summary Register
// @Description Register - API for registering users
// @Tags register
// @Accept  json
// @Produce  json
// @Param register body models.User true "register"
// @Success 200 {object} models.RegisterResponseModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /register [post]
func (h *handlerV1) Register(c *gin.Context) {
	var (
		body models.User
		code string
	)

	

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	err = body.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Message: err.Error(),
			},
		})
		h.log.Error("Error while validating", l.Error(err))
		return
	}

	// Checking uniqueness of username
	checkUsername, err := h.serviceManager.UserService().CheckField(
		context.Background(), &pb.Checkfild{
			Fildname: "username",
			Fild:     body.Username,
		},
	)
	if err != nil {
		fmt.Println(checkUsername)
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Message: err.Error(),
			},
		})
		h.log.Error("Error while checking uniquess", l.Error(err))
		return
	}

	if checkUsername.Message != "OK" {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Username already exists",
			},
		})
		return
	}

	checkEmail, err := h.serviceManager.UserService().CheckField(
		context.Background(), &pb.Checkfild{
			Fildname: "email",
			Fild:     body.Email,
		},
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Message: err.Error(),
			},
		})
		h.log.Error("Error while checking email uniquess", l.Error(err))
		return
	}

	if checkEmail.Message != "OK" {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Email already exists",
			},
		})
		return
	}

	code = etc.GenerateCode(7)
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while generatin code",
			},
		})
		return
	}

	_, err = h.serviceManager.EmailService().Send(
		context.Background(), &pbe.Email{
			Subject:    "Code for verification",
			Body:       code,
			Recipients: string(body.Email),
		},
	)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Message: err.Error(),
			},
		})
		h.log.Error("Error while sending verification code to user", l.Error(err))
		return
	}

	data := models.UserData{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Code:      code,
		Username:  body.Username,
		Password:  body.Password,
		Email:     body.Email,
	}

	bodyJSON, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while marshal user data for setting with ttl to redis",
			},
		})
		return
	}
	fmt.Println(string(bodyJSON))
	fmt.Println(code)
	err = h.inMemoryStorage.SetWithTTL(code, string(bodyJSON), 86400)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while setting with ttl to redis",
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.RegisterResponseModel{
		Message: "Verification code has been sent to your email, please check and verify",
	})
}

// Verify ...
// @Summary Verify
// @Description returns access token
// @Tags register
// @Accept  json
// @Produce  json
// @Param code path string true "code"
// @Success 200 {object} models.User
// @Failure 404 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /verify/{code} [post]
func (h *handlerV1) Verify(c *gin.Context) {
	code := c.Param("code")

	// Getting code from redis

	user := models.UserData{}
	fmt.Println(code)
	userJSON, err := redis.String(h.inMemoryStorage.Get(code))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	// Checking whether received code is valid
	if code != user.Code {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Message: "ErrorCodeInvalidCode",
			},
		})
		h.log.Error("verification failed", l.Error(err))
		return
	}

	user.ID = uuid.New().String()
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while generating new uuid for user",
			},
		})
		return
	}

	h.jwtHandler = token.JWTHandler{
		SigninKey: h.cfg.SigninKey,
		Sub:       user.ID,
		Iss:       "user",
		Role:      "authorized",
		Aud: []string{
			"nt",
		},
		Log: h.log,
	}

	// Creating access and refresh tokens
	accessTokenString, refreshTokenString, err := h.jwtHandler.GenerateAuthJWT()
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while generating tokens",
			},
		})
		return
	}
	fmt.Println(accessTokenString, refreshTokenString)
	// Creating hash of a password
	hashedPassword, err := etc.GeneratePasswordHash(user.Password)
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while generating hash for password",
			},
		})
		return
	}

	checkEmail, err := h.serviceManager.UserService().CheckField(
		context.Background(), &pb.Checkfild{
			Fildname: "email",
			Fild:     user.Email,
		},
	)
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while checking field",
			},
		})
		return
	}

	if checkEmail.Message != "OK" {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Mail already exists",
			},
		})
		return
	}

	// Creating new user
	newuser := &pb.UserRes{
		Id:           user.ID,
		Email:        user.Email,
		Password:     string(hashedPassword),
		AcsessToken:  accessTokenString,
		RefreshTaken: refreshTokenString,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
	}
	fmt.Println(newuser.Id)
	resUser, err := h.serviceManager.UserService().Create(context.Background(), newuser)
	if err != nil {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Message: "Error while create new user",
			},
		})
		return
	}

	c.JSON(http.StatusOK, &models.VerifyResponseModel{
		ID:           resUser.Id,
		AccessToken:  resUser.AcsessToken,
		RefreshToken: resUser.RefreshTaken,
	})
}
