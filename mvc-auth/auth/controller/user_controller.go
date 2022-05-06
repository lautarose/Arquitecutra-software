package controller

import (
	"Arquitecutra-software/mvc-auth/auth/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertUser(c *gin.Context) {
	var userDto dto.UserDto
	//using BindJson method to serialize body with struct

	if err := c.BindJSON(&userDto); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(userDto)
	c.JSON(http.StatusAccepted, &userDto)
}
