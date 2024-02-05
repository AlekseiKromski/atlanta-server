package tcp_consumer

import (
	"alekseikromski.com/atlanta/adapters/datapoints_parser"
	"alekseikromski.com/atlanta/core"
	"alekseikromski.com/atlanta/modules/storage"
	"fmt"
	"log"
	"net"
	"strings"
)

type ServerConfig struct {
	Address string
	BufSize int
}

func NewServerConfig(address string, bufSize int) *ServerConfig {
	return &ServerConfig{
		Address: address,
		BufSize: bufSize,
	}
}

type Server struct {
	EventBus chan string
	config   *ServerConfig
	listener net.Listener
	parser   *datapoints_parser.DataPointsParser
	storage  storage.Storage
}

func NewServer(conf *ServerConfig) *Server {
	return &Server{
		EventBus: make(chan string, 1),
		config:   conf,
		parser:   datapoints_parser.NewDataPointsParser(),
	}
}

func (s *Server) Start(notifyChannel chan struct{}, requirements map[string]core.Module) {
	// Load requirements
	storage, err := s.getStorageFromRequirement(requirements)
	if err != nil {
		log.Printf("TCP consumer: cannot start listener: %s", err)
		return
	}
	s.storage = storage

	listener, err := net.Listen("tcp", s.config.Address)
	if err != nil {
		log.Printf("TCP consumer: cannot start listener: %s", err)
		return
	}
	s.listener = listener

	// notify, that server started
	notifyChannel <- struct{}{}

	log.Println("TCP consumer: server started")

	for {
		// Accept incoming connections
		conn, err := s.listener.Accept()
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				continue // ignore
			}

			log.Printf("TCP consumer: %v", err)
			continue
		}
		log.Printf("TCP consumer: new connection %s", conn.RemoteAddr())

		// Handle client connection in a goroutine
		go s.handle(conn)
	}
}

func (s *Server) Stop() {
	// Close tcp listener
	s.listener.Close()
	log.Printf("TCP consumer: listener closed")

	//Close event bus
	close(s.EventBus)
	log.Printf("TCP consumer: event bus closed")
}

func (s *Server) Require() []string {
	return []string{
		"storage",
	}
}

func (s *Server) Signature() string {
	return "tcp_consumer"
}

func (s *Server) getStorageFromRequirement(requirements map[string]core.Module) (storage.Storage, error) {
	storage, ok := requirements["storage"].(storage.Storage)
	if !ok {
		return nil, fmt.Errorf("requiremnt list has wrong storage requirement")
	}

	return storage, nil
}

func (s *Server) handle(conn net.Conn) {
	buf := make([]byte, s.config.BufSize)
	count, err := conn.Read(buf)
	if err != nil {
		log.Printf("TCP consumer: cannot read message: %v", err)
		return
	}

	message := string(buf[:count])

	log.Printf("TCP consumer: received %s", message)

	if err := s.parser.Parse(message); err != nil {
		log.Printf("TCP consumer: cannot parse message: %v", err)
		return
	}

	if err := s.storage.SaveDatapoints(s.parser.Datapoints); err != nil {
		log.Printf("TCP consumer: cannot save datapoints: %v", err)
		return
	}

	s.EventBus <- message
}
