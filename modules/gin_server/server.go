package gin_server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
)

type ServerConfig struct {
	Port int
}

func NewServerConfig(port int) *ServerConfig {
	return &ServerConfig{
		Port: port,
	}
}

type Server struct {
	config *ServerConfig
	server *http.Server
}

func NewServer(conf *ServerConfig) *Server {
	return &Server{
		config: conf,
	}
}

func (s *Server) Start(notifyChannel chan struct{}) {
	log.Println("HTTP Server: init http server")

	router := gin.Default()

	// Register all handlers
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	// Create tcp listener and server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		log.Printf("HTTP Server: %s", err)
	}
	s.server = &http.Server{
		Handler: router,
	}

	// Notify core, that we started listener
	notifyChannel <- struct{}{}

	// Start server
	if err := s.server.Serve(listener); err != nil {
		log.Printf("HTTP Server: %s\n", err)
	}
}

func (s *Server) Stop() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP: cannot stop server: %s", err)
	}
}
