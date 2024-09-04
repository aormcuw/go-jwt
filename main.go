package main

import (
	"github.com/aormcuw/go-jwt/controllers"
	"github.com/aormcuw/go-jwt/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVars()
	initializer.ConnectToDB()
	initializer.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.Run()
}
