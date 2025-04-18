package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gomorpher/docs"
	"gomorpher/internal/api"
	"log"
	"net/http"
)

// @title Gomorpher API
// @version 1.0
// @description API

// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API works!",
		})
	})

	apiGroup := r.Group("/api/v1")
	apiGroup.POST("/generate", api.GenerateHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
