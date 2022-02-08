package handlers

import (
	"context"
	"net/http"
	"time"

	pb "Golang/apiGateway/genproto"
	l "Golang/apiGateway/pkg/logger"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateTask godoc
// @Summary Create new task
// @Schemes
// @Description  Create new task
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} task.Task
// @Failure 400 {object} models.StandartErrorModel
// @Failure 500 {object} models.StandartErrorModel
// @Router /v1/users [post]
func (h *handlerV1) CreateTask(c *gin.Context) {
	var (
		body        pb.Task
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

	response, err := h.serviceManager.TaskService().Create(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create task", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetTask godoc
// @Summary Get task
// @Schemes
// @Description  Get task
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} task.Task
// @Router /users/{id} [get]
func (h *handlerV1) GetTask(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().GetTask(
		ctx, &pb.Task{
			Id: guid,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get task", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Update task
// @Schemes
// @Description  Update task
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} task.Task
// @Router /users [put]
func (h *handlerV1) UpdateTask(c *gin.Context) {
	var (
		body        pb.Task
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

	response, err := h.serviceManager.TaskService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Delete task
// @Schemes
// @Description  Delete task
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} task.Mess
// @Router /users/{id} [delete]
func (h *handlerV1) DeleteTask(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Delete(
		ctx, &pb.Task{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete task", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
