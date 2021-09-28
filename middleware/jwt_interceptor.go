package middleware

import (
	"fmt"
	"net/http"
	"os"
	"stock/model"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(user model.User) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["id"] = user.ID
	atClaims["username"] = user.Username
	atClaims["password"] = user.Password
	atClaims["level"] = user.Level
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return " ", err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	tkArr := strings.Split(bearToken, " ")
	if len(tkArr) == 2 {
		return tkArr[1]
	}
	return " "
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexperted siging method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func AuthMiddleware(c *gin.Context) {
	token, err := VerifyToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": err})
		c.Abort()
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("jwt_username", claims["username"])
		c.Set("jwt_level", claims["level"])
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Invalid token"})
		c.Abort()
	}
}
