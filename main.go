package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/send2gopal/gocrud/database"
	"github.com/send2gopal/gocrud/database/models"
	"github.com/send2gopal/gocrud/initializer"
)

func init() {
	initializer.LoadEnvironmentvariables()
	database.ConnectToDB()
	database.DB.AutoMigrate(&models.User{})
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
