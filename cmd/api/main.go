package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := server(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func server() error {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return router.Run()
}
