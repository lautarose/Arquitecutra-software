package controller

import (
	"Arquitecutra-software/mvc-auth/mvc_ej_gingonic/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertBody(c *gin.Context) {
	var body dto.BodyDto
	//using BindJson method to serialize body with struct

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(body)
	c.JSON(201, &body)
}
