package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hieuocb/go-gin-gorm-rest/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetAllUsers)
	router.GET("/:id", controller.GetUser)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
}
