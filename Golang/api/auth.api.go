package api

import (
	"rest/db"
	"rest/interceptor"
	"rest/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// SetupAuthenAPI - OK
func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v1")
	{
		authenAPI.POST("/login", login)
	}
}

func login(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		var queryUser model.User
		if err := db.GetDatabase().First(&queryUser, "email = ?", user.Email).Error; err != nil {
			c.JSON(200, gin.H{"result": "nok", "error": err})
		} else if checkPasswordHash(user.Password, queryUser.Password) == false {
			c.JSON(200, gin.H{"result": "nok", "error": "invalid password"})
		} else {
			token := interceptor.JwtSign(queryUser)
			c.JSON(200, gin.H{"result": "ok", "token": token})
		}

	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
