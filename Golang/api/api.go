package api

import (
	"rest/db"

	"github.com/gin-gonic/gin"
)

// Setup - call this
func Setup(router *gin.Engine) {
	db.ConnectDB()
	SetupAuthenAPI(router)
}
