package tcp_consumer

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type ServerConfig struct {
	Port    int
	BufSize int
}

func NewServerConfig(port, bufSize int) *ServerConfig {
	return &ServerConfig{
		Port:    port,
		BufSize: bufSize,
	}
}

type Server struct {
	EventBus chan string
	config   *ServerConfig
	listener net.Listener
}

func NewServer(conf *ServerConfig) *Server {
	return &Server{
		EventBus: make(chan string, 1),
		config:   conf,
	}
}

func (s *Server) Start(notifyChannel chan struct{}) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		log.Printf("TCP consumer: %s", err)
	}
	s.listener = listener

	// notify, that server started
	notifyChannel <- struct{}{}

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
	if err := s.listener.Close(); err != nil {
		log.Printf("TCP consumer: cannot read message")
	}

	log.Printf("TCP consumer: listener closed")

	//Close event bus
	close(s.EventBus)
	log.Printf("TCP consumer: event bus closed")
}

func (s *Server) handle(conn net.Conn) {
	buf := make([]byte, s.config.BufSize)
	count, err := conn.Read(buf)
	if err != nil {
		log.Printf("TCP consumer: cannot read message")
		return
	}

	message := string(buf[:count])

	s.EventBus <- message

	// TODO: parse and save
	//parser := datapoints_parser.NewDataPointsParser(message)
	log.Printf("TCP consumer: received %s", message)
}
