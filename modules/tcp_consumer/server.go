package tcp_consumer

import (
	"alekseikromski.com/atlanta/adapters/datapoints_parser"
	"alekseikromski.com/atlanta/core"
	"alekseikromski.com/atlanta/modules/storage"
	"log"
	"net"
)

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

	s.listen()
}

func (s *Server) Stop() {
	// Close tcp listener
	s.listener.Close()
	log.Printf("TCP consumer: listener closed")

	//Close event bus
	close(s.EventBus)
	log.Printf("TCP consumer: event bus closed")
}

func (s *Server) Signature() string {
	return "tcp_consumer"
}
