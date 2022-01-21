package handlers

import (
	"app/database/custumer"
	"app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Customer(c *gin.Context) {
	var costumer models.Costumer
	c.BindJSON(&costumer)
	 err := custumer.Create(costumer)
	if err != nil {
		panic(err)
	}
	c.JSON(201,costumer)
}

func Login(c *gin.Context) {
	var costumer models.Costumer
	c.BindJSON(&costumer)
	 new,err := custumer.Login(costumer)
	if err != nil {
		panic(err)
	}
	c.JSON(201,new)
}

func Query(c *gin.Context) {
	var costumer models.Costumer
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		panic(err)
	}
	costumer.Id = int64(id)
	 new,err := custumer.Query_id(costumer)
	if err != nil {
		panic(err)
	}
	c.JSON(201,new)
}


