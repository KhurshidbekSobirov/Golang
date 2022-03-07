package handlers

import (
	"context"
	"fmt"
	pb "myProject/apiGatwey/genproto/user_service"
	l "myProject/apiGatwey/pkg/logger"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateUser godoc
// @Summary Create new user
// @Description This API for creating a new user
// @Tags User
// @Accept json
// @Param body body user.UserRes true "body"
// @Produce json
// @Success 201 {object} user.UserReq
// @Router /users [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		body        pb.UserRes
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	fild := body.Email
	mes, err := h.serviceManager.UserService().CheckField(ctx, &pb.Checkfild{
		Fildname: "email",
		Fild:     fild,
	})

	if mes.Message != "OK" || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exists",
		})
	} else {
		fild = body.Username
		mes, err = h.serviceManager.UserService().CheckField(ctx, &pb.Checkfild{
			Fildname: "username",
			Fild:     fild,
		})

		if mes.Message != "OK" || err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username already exists",
			})
		} else {
			fmt.Println(mes)
			response, err := h.serviceManager.UserService().Create(ctx, &body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				h.log.Error("failed to create user", l.Error(err))
				return
			}

			c.JSON(http.StatusCreated, response)
		}
	}

}

// GetUser godoc
// @Summary GetUser
// @Schemes
// @Description  Get User
// @Tags User
// @Accept json
// @Param id path string true "ID"
// @Produce json
// @Success 200 {object} user.UserReq
// @Router /user/{id} [get]
func (h *handlerV1) GetUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetUser(
		ctx, &pb.ById{
			UserId: guid,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get User", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)

}

// @Summary Update User
// @Schemes
// @Description  Update User
// @Tags User
// @Accept json
// @Param body body user.UserReq true "body"
// @Produce json
// @Success 200 {object} user.Mess
// @Router /user/{id} [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		body        pb.UserReq
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	fild := body.Email
	mes, err := h.serviceManager.UserService().CheckField(ctx, &pb.Checkfild{
		Fildname: "email",
		Fild:     fild,
	})

	if mes.Message != "OK" || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exists",
		})
	} else {
		fild = body.Username
		mes, err = h.serviceManager.UserService().CheckField(ctx, &pb.Checkfild{
			Fildname: "username",
			Fild:     fild,
		})

		if mes.Message != "OK" || err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username already exists",
			})
		} else {
			response, err := h.serviceManager.UserService().Update(ctx, &body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				h.log.Error("failed to update user", l.Error(err))
			}

			c.JSON(http.StatusOK, response)
		}
	}

}

// @Summary DeleteUser
// @Schemes
// @Description  Delete User
// @Tags User
// @Param body body user.ById true "body"
// @Accept json
// @Produce json
// @Success 200 {object} user.Mess
// @Router /user/{id} [delete]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Delete(
		ctx, &pb.ById{
			UserId: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete User", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary GetMyProfile
// @Security BearerAuth
// @Description GetMyProfile
// @Tags User
// @Produce json
// @Success 200 {object} user.UserReq
// @Failure 404 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /profile [get]
func (h *handlerV1) GetMyProfile(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

    claims := GetClaims(h ,c )

	userId := claims["sub"].(string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetUser(
		ctx, &pb.ById{
			UserId: userId,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get User", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)


}

// @Summary Login
// @Schemes
// @Description  Login
// @Tags register
// @Accept json
// @Param body body user.GetByemail true "Email"
// @Produce json
// @Success 200 {object} user.UserReq
// @Router /login [put]
func (h *handlerV1) LogIn(c *gin.Context){

	var (
		body        pb.UserReq
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetByEmail(
		ctx, &pb.GetByemail{
			Email: body.Email,
			Password: body.Password,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get User", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
	
}

