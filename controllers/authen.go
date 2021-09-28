package controllers

import (
	"net/http"
	"stock/database"
	"stock/middleware"
	"stock/model"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *gin.Context) {
	var user model.User
	var queryUser model.User
	db := database.GetDB()
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Invalid json provided"})
		return
	}
	if err := db.First(&queryUser, "username = ?", user.Username).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "invalid username"})
		return
	}
	if checkPasswordHash(user.Password, queryUser.Password) == false {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "invalid password"})
		return
	}
	token, err := middleware.CreateToken(queryUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Register(c *gin.Context) {
	var user model.User
	var err error
	db := database.GetDB()
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Invalid json provided"})
		return
	}
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": err})
		return
	}
	user.CreatedAt = time.Now()
	db.AutoMigrate(&model.User{})
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
