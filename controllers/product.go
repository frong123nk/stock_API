package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"stock/database"
	"stock/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {

}

func GetProduct(c *gin.Context) {
	var product []model.Product
	db := database.GetDB()
	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		err := db.Raw("select * from products where name like ?", keyword).Scan(&product).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "Not Found Data"})
			return
		}
	} else {
		err := db.Find(&product).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "Not Found Data"})
			return
		}
	}
	c.JSON(http.StatusOK, product)
}

func flieExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func saveImage(image *multipart.FileHeader, product *model.Product, c *gin.Context) {
	db := database.GetDB()
	if image != nil {
		runningDir, _ := os.Getwd()
		product.Image = image.Filename
		extension := filepath.Ext(image.Filename)
		filename := fmt.Sprintf("%d%s", product.ID, extension)
		filePath := fmt.Sprintf("%s/uploaded/images/%s", runningDir, filename)
		if flieExists(filePath) {
			os.Remove(filePath)
		}
		c.SaveUploadedFile(image, filePath)
		db.Exec("update products set image = ? WHERE id = ?", filename, product.ID)
	}
}

func CreateProduct(c *gin.Context) {
	var product model.Product
	db := database.GetDB()
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	image, _ := c.FormFile("image")
	product.Image = image.Filename
	product.CreatedAt = time.Now()
	errdb := db.AutoMigrate(&model.Product{}).Error

	if errdb != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Can't connect database"})
		return
	}
	err := db.Create(&product).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not Found Database"})
		return
	}
	saveImage(image, &product, c)

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func EditProduct(c *gin.Context) {
	var product model.Product
	db := database.GetDB()
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	product.ID = uint(id)
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	image, _ := c.FormFile("image")

	err := db.Exec("update products set name =?,stock=?,price =? where id =?",
		product.Name, product.Stock, product.Price, product.ID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not Found Database"})
		return
	}
	saveImage(image, &product, c)
	c.JSON(http.StatusOK, gin.H{"product": product})

}

func DeleteProduct(c *gin.Context) {
	var product model.Product
	db := database.GetDB()
	product.Name = c.PostForm("name")
	err := db.Exec("delete from products where name like ?", product.Name).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not Found Data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Successfully deleted"})
}
