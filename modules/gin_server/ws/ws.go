package ws

import (
	"alekseikromski.com/atlanta/modules/gin_server/guard"
	"fmt"
	"github.com/AlekseiKromski/at-socket-server/core"
	"github.com/gin-gonic/gin"
	"log"
)

type WebSocket struct {
	app *core.App
}

func NewWebSocket(engine *gin.Engine, secret []byte, guard *guard.Guard) (*WebSocket, error) {
	app, err := core.Start(
		engine,
		&core.Handlers{},
		guard.Check,
		&core.Config{
			JwtSecret: secret,
			Debug:     true,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("cannot start websocket server: %v", err)
	}

	go func() {
		for {
			event := <-app.Hooks
			log.Println(event.Data)
		}
	}()

	return &WebSocket{
		app: app,
	}, nil
}

func (ws *WebSocket) SendDatapointsToAllClients(data string) error {
	for _, client := range ws.app.Clients {
		log.Println(client.Send(data, "INFO"))
	}
	return nil
}
