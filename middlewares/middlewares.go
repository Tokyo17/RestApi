package middlewares

import (
	// "fmt"
	jwt "github.com/golang-jwt/jwt"
	"restApi/contants"
	"time"
)

type JwtCustomClaims struct {
	UserId int `json:"name"`
	jwt.StandardClaims
}

func CreateToken(userId int) (string, error) {
	// fmt.Println(userId)
	claims := &JwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(contants.SECRET_JWT))

}
