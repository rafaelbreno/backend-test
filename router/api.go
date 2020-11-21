package router

import (
	"coderockr-test/app/controllers"
	"coderockr-test/app/models"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Init() {
	r = gin.Default()

	models.ConnectDatabase()

	auth()

	r.Run(":8080")
}

func auth() {
	c := controllers.UserController{}

	r.POST("/signup", c.SignUp)
}
