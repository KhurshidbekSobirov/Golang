package main

import (
	"app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/costumer",handlers.Customer)
	r.GET("/costumer/login",handlers.Login)
	r.GET("/costumer/:id",handlers.Query)
	r.Run()
}