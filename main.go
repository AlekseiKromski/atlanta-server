package main

import (
	"alekseikromski.com/atlanta/server"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

//go:embed examples
var examplesFS embed.FS

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("cannot load env form .env file: %v", err)
		return

	}
	port := os.Getenv("PORT")
	udpPort := os.Getenv("UDP_PORT")

	r := gin.Default()
	r.GET("/:file", func(context *gin.Context) {

		file, err := examplesFS.ReadFile(fmt.Sprintf("examples/%s", context.Param("file")))
		if err != nil {
			context.JSON(http.StatusInternalServerError, "cannot find file, context to admin")
			log.Printf("cannot read file: %v", err)
			return
		}

		if _, err := context.Writer.Write(file); err != nil {
			context.JSON(http.StatusInternalServerError, "cannot find file, context to admin")
			log.Printf("cannot write file: %v", err)
			return
		}

		context.Status(http.StatusOK)
		context.Header("Content-Type", "application/json")
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	//Start upd server
	udpServer := server.NewUDPServer(udpPort)
	go func() {
		err := udpServer.Start()
		if err != nil {
			log.Printf("Cannot run UDP server: %v", err)
			return
		}
	}()

	//Start API server
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Printf("error during running server: %v", err)
	}
}
