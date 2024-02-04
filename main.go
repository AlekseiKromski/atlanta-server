package main

import (
	"alekseikromski.com/atlanta/core"
	"alekseikromski.com/atlanta/modules/gin_server"
	"alekseikromski.com/atlanta/modules/tcp_consumer"
)

func main() {
	c := core.NewCore()
	c.Init([]core.Module{
		gin_server.NewServer(
			gin_server.NewServerConfig(3000), // main application port
		),
		tcp_consumer.NewServer(
			tcp_consumer.NewServerConfig(3001, 250), // Datapoints port
		),
	})
}
