package server

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"log"
	"net"
	"testing"
)

func TestUdpConnection(t *testing.T) {
	port := 3000

	//start app
	udp := NewUDPServer(fmt.Sprintf("%d", port))
	if err := udp.Start(); err != nil {
		t.Fatalf("Cannot run UDP Server: %v", err)
		return
	}

	defer udp.Server.Close()

	<-udp.Ping

	udpServer, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))

	if err != nil {
		t.Fatalf("ResolveUDPAddr failed: %v", err.Error())
		return
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		t.Fatalf("Listen failed: %v", err.Error())
		return
	}

	_, err = conn.Write([]byte("Hello world"))
	if err != nil {
		t.Fatalf("Cannot write message: %v", err.Error())
		return
	}

	for {
		response := make([]byte, 2)
		_, addr, err := conn.ReadFromUDP(response)
		if err != nil {
			t.Fatalf("Cannot read message: %v", err.Error())
			return
		}

		log.Printf("Client received [%s]: %s", addr.IP, response)

		assert.Equal(t, string(response), "OK")
		conn.Close()
		break
	}
}
