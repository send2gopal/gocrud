package main

import (
	"github.com/gin-gonic/gin"
	"github.com/send2gopal/gocrud/controllers"
	"github.com/send2gopal/gocrud/database"
	"github.com/send2gopal/gocrud/initializer"
)

func init() {
	initializer.LoadEnvironmentvariables()
	database.ConnectToDB()
	database.MigrateDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/users", controllers.CreateUser)
	r.Run()
}
