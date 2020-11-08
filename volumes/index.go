package main

import (
	"fmt"
	"os"
	"rest/api"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	runningDir, _ := os.Getwd()
	errlogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	gin.DefaultErrorWriter = errlogfile
	gin.DefaultWriter = accesslogfile

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s  \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}

	router.Use(cors.New(config))

	router.Static("/uploads", "/uploads")

	api.Setup(router)

	fmt.Println("Environment Port : 8001")
	router.Run(":8001")
}
