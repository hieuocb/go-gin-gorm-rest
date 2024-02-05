package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hieuocb/go-gin-gorm-rest/config"
	"github.com/hieuocb/go-gin-gorm-rest/routes"
)

func main() {
	router := gin.Default()
	cfg := config.LoadConfig()
	if _, err := config.Connect(cfg.DatabaseURI); err != nil {
		return
	}

	//router.Use(middleware.ErrorHandler)
	routes.PingRoute(router)

	routes.UserRoute(router)
	addr := "localhost:" + cfg.HttpPort
	if err := router.Run(addr); err != nil {
		return
	}
}
