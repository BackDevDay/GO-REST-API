package interceptor

import (
	"rest/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "dddsdfsaf"

// JwtSign -
func JwtSign(payload model.User) string {
	atClaims := jwt.MapClaims{}

	// Payload begin
	atClaims["id"] = payload.ID
	atClaims["username"] = payload.Email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// Payload end

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secretKey))
	return token

}
