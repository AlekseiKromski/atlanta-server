package tcp_consumer

import (
	"alekseikromski.com/atlanta/core"
	"alekseikromski.com/atlanta/modules/storage/postgres"
	"fmt"
	"net"
	"testing"
)

func TestConnection(t *testing.T) {
	notifyChannel := make(chan struct{}, 1)
	postgres := postgres.NewPostgres(
		postgres.NewConfig(
			"localhost",
			"atlanta",
			"postgres",
			"postgres",
			5432,
		),
	)
	go postgres.Start(notifyChannel, map[string]core.Module{})
	<-notifyChannel
	defer postgres.Stop()

	port := 3001

	// Start tcp server
	tcpServer := NewServer(
		NewServerConfig(fmt.Sprintf(":%d", port), 250),
	)

	// Start & wait server
	go tcpServer.Start(notifyChannel, map[string]core.Module{
		"storage": postgres,
	})
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
	content := "DEVICE::3cc76ff4-cbaa-436c-b727-45d526facfc7;TIME::2019-10-12T07:20:50.52Z;TEMP::14;PRS::1000PA;HUM::32.0"
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
