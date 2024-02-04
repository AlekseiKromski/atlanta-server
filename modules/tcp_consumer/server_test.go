package tcp_consumer

import (
	"fmt"
	"net"
	"testing"
)

func TestConnection(t *testing.T) {
	port := 3001
	notifyChannel := make(chan struct{}, 1)

	// Start tcp server
	tcpServer := NewServer(
		NewServerConfig(port, 250),
	)

	// Start & wait server
	go tcpServer.Start(notifyChannel)
	<-notifyChannel
	defer tcpServer.Stop()

	// Create client
	clientConn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
	defer clientConn.Close()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Send data to tcp server
	content := "tcp content"
	if _, err = clientConn.Write([]byte(content)); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Check received server content by event bus
	if content != <-tcpServer.EventBus {
		t.Fatalf("server received unexpected data: %s", content)
		return
	}

}
