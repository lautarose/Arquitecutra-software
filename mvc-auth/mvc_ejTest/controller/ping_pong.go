package controller

import "github.com/gin-gonic/gin"

func PingPong(c *gin.Context) {
	c.String(200, "pong")
}
