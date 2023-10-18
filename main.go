package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("cannot load env form .env file: %v", err)
		return
	}
	port := os.Getenv("PORT")

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Printf("error during running server: %v", err)
	}
}
