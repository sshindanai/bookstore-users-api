package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sshindanai/repo/bookstore-users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapURLs()

	logger.Info("app is started...")
	router.Run(":8080")
}
