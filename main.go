package main

import (
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
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.Run()
}
