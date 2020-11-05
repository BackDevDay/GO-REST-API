package api

import (
	"go-rest-api/db"

	"github.com/gin-gonic/gin"
)

// Setup - call this method to setup routes
func Setup(router *gin.Engine) {

	db.ConnectDB()
}