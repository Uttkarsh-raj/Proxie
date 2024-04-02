package routes

import (
	"github.com/Uttkarsh-Raj/Proxie/controller"
	"github.com/gin-gonic/gin"
)

func IncomingRoutes(router *gin.Engine) {
	router.GET("/", controller.ProxyServer())
}
