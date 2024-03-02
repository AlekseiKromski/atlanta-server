package tcp_consumer

import (
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

			s.Log("cannot accept connection", err.Error())
			continue
		}
		s.Log("new connection", conn.RemoteAddr().String())

		// Handle client connection in a goroutine
		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	buf := make([]byte, s.config.bufSize)
	count, err := conn.Read(buf)
	if err != nil {
		s.Log("cannot read message", err.Error())
		return
	}

	message := string(buf[:count])

	s.Log("received", message)

	deviceUuid, datapoints, err := s.parser.Parse(message)
	if err != nil {
		s.Log("cannot parse message: ", err.Error(), message)
		return
	}

	if len(deviceUuid) == 0 {
		s.Log("empty device id, ignored", message)
		return
	}

	if err := s.storage.SaveDatapoints(deviceUuid, datapoints); err != nil {
		s.Log("cannot save datapoints", err.Error())
		return
	}

	s.EventBus <- message
}
