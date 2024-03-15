package tcp_consumer

import (
	"alekseikromski.com/atlanta/adapters/datapoints_parser"
	"alekseikromski.com/atlanta/core"
	"alekseikromski.com/atlanta/modules/storage"
	"net"
)

type Server struct {
	EventBus        chan string
	config          *ServerConfig
	listener        net.Listener
	parser          *datapoints_parser.DataPointsParser
	storage         storage.Storage
	eventBusChannel chan core.BusEvent
}

func NewServer(conf *ServerConfig) *Server {
	return &Server{
		EventBus: make(chan string, 1),
		config:   conf,
		parser:   datapoints_parser.NewDataPointsParser(),
	}
}

func (s *Server) Start(notifyChannel chan struct{}, eventBusChannel chan core.BusEvent, requirements map[string]core.Module) {
	s.eventBusChannel = eventBusChannel

	// Load requirements
	storage, err := s.getStorageFromRequirement(requirements)
	if err != nil {
		s.Log("cannot start listener", err.Error())
		return
	}
	s.storage = storage

	listener, err := net.Listen("tcp", s.config.address)
	if err != nil {
		s.Log("cannot start listener", err.Error())
		return
	}
	s.listener = listener

	// notify, that server started
	notifyChannel <- struct{}{}

	s.Log("server started")

	s.listen()
}

func (s *Server) Stop() {
	// Close tcp listener
	s.listener.Close()

	s.Log("listener closed")

	//Close event bus
	close(s.EventBus)
	s.Log("event bus closed")
}

func (s *Server) Signature() string {
	return "tcp_consumer"
}
