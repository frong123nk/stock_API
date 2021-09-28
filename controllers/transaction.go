package controllers

import (
	"fmt"
	"net/http"
	"stock/database"
	"stock/model"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTransaction(c *gin.Context) {
	var transaction []model.Transaction
	db := database.GetDB()
	staff_name := c.GetString("jwt_username")
	err := db.Raw("select * from transactions where staff_name like ?", staff_name).Scan(&transaction).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not Found Data"})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func CreateTransaction(c *gin.Context) {
	var transaction model.Transaction
	db := database.GetDB()
	db.AutoMigrate(&transaction)
	transaction.StaffName = c.GetString("jwt_username")
	fmt.Println("Staff_name", transaction.StaffName)
	transaction.CreatedAt = time.Now()
	if c.ShouldBind(&transaction) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Invalid json provided"})
		return
	}
	transaction.Change = transaction.Paid - transaction.Total
	if transaction.Change > 0 {
		transaction.PaymentDetail = "FULL"
	} else {
		transaction.PaymentDetail = "MINUS"
	}
	if err := db.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}
