package server

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

type Udp struct {
	port string

	Server *net.UDPConn
	Ping   chan struct{}
}

func NewUDPServer(port string) *Udp {
	return &Udp{
		port: port,
		Ping: make(chan struct{}),
	}
}

func (u *Udp) Start() error {
	server, err := net.ResolveUDPAddr("udp4", fmt.Sprintf(":%s", u.port))
	if err != nil {
		return fmt.Errorf("cannot resolve UDP address Server: %v", err)
	}

	connection, err := net.ListenUDP("udp4", server)
	if err != nil {
		return fmt.Errorf("cannot start UDP Server: %v", err)
	}

	u.Server = connection

	log.Printf("Start UDP Server")

	go u.run()

	return nil
}

func (u *Udp) run() {
	u.Ping <- struct{}{}
	for {
		buf := make([]byte, 1024)
		_, addr, _ := u.Server.ReadFromUDP(buf)

		if len(buf) == 0 || addr == nil {
			continue
		}

		data := bytes.Trim(buf, "\x00")
		log.Printf("Server received [%s]: %s", addr, data)

		if _, err := u.Server.WriteTo([]byte("OK"), addr); err != nil {
			log.Printf("UDP Server error while writing [%s]: %v", addr.IP, err)
			continue
		}
	}
}
