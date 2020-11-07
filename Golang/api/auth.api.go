package api

import (
	"net/http"
	"rest/db"
	"rest/interceptor"
	"rest/model"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// SetupAuthenAPI - OK
func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v1")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

func login(context *gin.Context) {
	var user model.User
	if context.ShouldBind(&user) == nil {
		var queryUser model.User
		if err := db.GetDatabase().First(&queryUser, "email = ?", user.Email).Error; err != nil {
			context.JSON(http.StatusOK, gin.H{"status": false, "message": err})
		} else if checkPasswordHash(user.Password, queryUser.Password) == false {
			context.JSON(http.StatusOK, gin.H{"status": false, "message": "invalid password"})
		} else {
			token := interceptor.JwtSign(queryUser)
			context.JSON(http.StatusOK, gin.H{"status": true, "message": "ok", "token": token})
		}

	} else {
		context.JSON(http.StatusNotFound, gin.H{"status": false, "message": "unable to bind data"})
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func register(context *gin.Context) {
	var user model.User
	if context.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		user.CreatedAt = time.Now()
		if err := db.GetDatabase().Create(&user).Error; err != nil {
			context.JSON(http.StatusOK, gin.H{"status": true, "message": err})
		} else {
			context.JSON(http.StatusOK, gin.H{"status": true, "message": "ok", "data": user})
		}
	} else {
		context.JSON(http.StatusNotFound, gin.H{"status": false, "message": "unable to bind data"})
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
