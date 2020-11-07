package main

import (
	"fmt"
	"rest/api"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Static("/uploads", "../volumes/uploads")

	api.Setup(router)

	fmt.Println("Environment Port : 8001")
	router.Run(":8001")
}
