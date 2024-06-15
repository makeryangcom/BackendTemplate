package Router

import (
	"Platform/framework/config"
	"Platform/framework/controller/ping"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	router := gin.New()

	gin.SetMode(Config.Get.Service.Mode)

	ping := router.Group("ping")
	{
		ping.GET("/index", PingController.Ping)
	}

	return router
}
