package main

import (
	"fmt"
	"go-rest-api/api"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Static("/uploads", "./uploads")

	api.Setup(router)

	fmt.Println("Environment Port : 9001")
	router.Run(":9001")
}
