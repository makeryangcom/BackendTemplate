package PingController

import (
	"Platform/framework/utils"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	Utils.Success(c, Utils.EmptyData{})
	return
}
