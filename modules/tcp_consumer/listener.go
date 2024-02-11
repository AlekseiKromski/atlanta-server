package tcp_consumer

import (
	"log"
	"net"
	"strings"
)

func (s *Server) listen() {
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
