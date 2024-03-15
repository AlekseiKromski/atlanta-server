package main

import (
	"alekseikromski.com/atlanta/core"
	"alekseikromski.com/atlanta/modules/gin_server"
	"alekseikromski.com/atlanta/modules/storage/postgres"
	"alekseikromski.com/atlanta/modules/tcp_consumer"
	"embed"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	//go:embed front-end/build
	resources embed.FS
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Main: cannot load env form .env file: %v", err)
		return
	}

	ginAddress := os.Getenv("GIN_ADDRESS")
	ginSecret := os.Getenv("GIN_SECRET")
	tcpConsumerAddress := os.Getenv("TCP_CONSUMER_ADDRESS")

	tcpConsumerBuf, err := strconv.Atoi(os.Getenv("TCP_CONSUMER_BUF"))
	if err != nil {
		log.Printf("Main: cannot load TCP_CONSUMER_BUF: %v", err)
		return
	}

	dbDatabase := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Printf("Main: cannot load DB_PORT: %v", err)
		return
	}

	c := core.NewCore()
	c.Init([]core.Module{
		gin_server.NewServer(
			gin_server.NewServerConfig(ginSecret, ginAddress),
			resources,
		),
		tcp_consumer.NewServer(
			tcp_consumer.NewServerConfig(tcpConsumerAddress, tcpConsumerBuf),
		),
		postgres.NewPostgres(
			postgres.NewConfig(
				dbHost,
				dbDatabase,
				dbUsername,
				dbPassword,
				dbPort,
			),
		),
	})
}
