package gin_server

import (
	"alekseikromski.com/atlanta/core"
	v1 "alekseikromski.com/atlanta/modules/gin_server/v1"
	"context"
	"log"
	"net"
	"net/http"
)

type ServerConfig struct {
	Address string
}

func NewServerConfig(address string) *ServerConfig {
	return &ServerConfig{
		Address: address,
	}
}

type Server struct {
	config *ServerConfig
	server *http.Server
	api    Api
}

func NewServer(conf *ServerConfig) *Server {
	return &Server{
		config: conf,
	}
}

func (s *Server) Start(notifyChannel chan struct{}, requirements map[string]core.Module) {
	log.Println("HTTP Server: init http server")

	storage, err := s.getStorageFromRequirement(requirements)
	if err != nil {
		log.Printf("HTTP Server: %s", err)
		return
	}

	s.api = v1.NewV1Api(storage, s.Log)

	if err := s.api.RegisterRoutes(); err != nil {
		log.Printf("HTTP Server: %s", err)
		return
	}

	// Create tcp listener and server
	listener, err := net.Listen("tcp", s.config.Address)
	if err != nil {
		log.Printf("HTTP Server: %s", err)
		return
	}
	s.server = &http.Server{
		Handler: s.api.GetEngine(),
	}

	// Notify core, that we started listener
	notifyChannel <- struct{}{}

	// Start server
	if err := s.server.Serve(listener); err != nil {
		log.Printf("HTTP Server: %s\n", err)
		return
	}
}

func (s *Server) Stop() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP: cannot stop server: %s", err)
		return
	}
}

func (s *Server) Signature() string {
	return "gin_server"
}
