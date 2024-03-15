package tcp_consumer

import (
	"alekseikromski.com/atlanta/core"
	"alekseikromski.com/atlanta/modules/storage/postgres"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	notifyChannel := make(chan struct{}, 1)
	busEventChannel := make(chan core.BusEvent, 1)
	postgres := postgres.NewPostgres(
		postgres.NewConfig(
			"localhost",
			"atlanta",
			"postgres",
			"postgres",
			5432,
		),
	)

	go postgres.Start(notifyChannel, busEventChannel, map[string]core.Module{})
	<-notifyChannel
	defer postgres.Stop()

	port := 3001

	// Start tcp server
	tcpServer := NewServer(
		NewServerConfig(fmt.Sprintf(":%d", port), 250),
	)

	// Start & wait server
	go tcpServer.Start(notifyChannel, busEventChannel, map[string]core.Module{
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
	content := "DEVICE::3cc76ff4-cbaa-436c-b727-45d526facfc7;TIME::2019-10-12T07:20:50.52Z;TEMP::14;PRS::1000;HUM::32.0;GEO::35.22,35.22"
	if _, err = clientConn.Write([]byte(content)); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Check received server content by event bus
	if content != <-tcpServer.EventBus {
		t.Fatalf("server received unexpected data: %s", content)
		return
	}

	event := <-tcpServer.eventBusChannel
	assert.Equal(t, "gin_server", event.Receiver)
}

func Test(t *testing.T) {
	for {
		conn, err := net.Dial("tcp", "localhost:3017")
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}
		_, err = conn.Write([]byte("DEVICE::3cc76ff4-cbaa-436c-b727-45d526facfc7;HUM::43.00;TEMP::18.00;PRS::101738;ALT2::-34.91;TEMP2::16.50;GEO::59.337040,27.420391;TIME::" + time.Now().Format(time.RFC3339) + ";ALT::5.00"))
		if err != nil {
			log.Println(err)
		}
		conn.Close()
		time.Sleep(5 * time.Second)
	}
}
