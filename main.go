package main

import (
	"github.com/gin-gonic/gin"
	clientController "github.com/sahildotexe/go-rate-limit/controllers"
	"github.com/sahildotexe/go-rate-limit/middlewares"
)

func main() {
	r := gin.Default()
	client := r.Group("/token")
	client.GET("/", clientController.GenerateClientKey)
	endpoint := r.Group("/ping")
	endpoint.Use(middlewares.Limit)
	endpoint.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()

}
