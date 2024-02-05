package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hieuocb/go-gin-gorm-rest/controller"
)

func PingRoute(router *gin.Engine) {
	router.GET("/ping", controller.PingController)
}
