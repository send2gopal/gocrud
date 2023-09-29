package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/send2gopal/gocrud/database"
	"github.com/send2gopal/gocrud/database/models"
)

func CreateUser(c *gin.Context) {

	var user models.User

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	result := database.DB.Create(&user)

	if result.Error != nil {
		log.Fatal("Error while creating user.", result.Error)
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
