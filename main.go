package main

import (
	"fmt"
	"stock/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
		return
	}
	router := gin.Default()
	// router.Static("/images", "./uploaded/images")
	routers.Setup(router)
	router.Run(":85")
}
