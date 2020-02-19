package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vSivarajah/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	MapUrls()
	logger.Info("About to start the application")
	router.Run(":8080")
}
