package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jooyungik/auth-server/middlewares"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.InjectDB(InitDB()))
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

}
