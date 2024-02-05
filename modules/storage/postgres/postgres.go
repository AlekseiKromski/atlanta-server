package postgres

import (
	"alekseikromski.com/atlanta/core"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Config struct {
	host     string
	port     int
	database string
	username string
	password string
}

func NewConfig(host, database, username, password string, port int) *Config {
	return &Config{
		host:     host,
		port:     port,
		database: database,
		username: username,
		password: password,
	}
}

type Postgres struct {
	config *Config
	db     *sql.DB
}

func NewPostgres(config *Config) *Postgres {
	return &Postgres{
		config: config,
		db:     nil,
	}
}

func (p *Postgres) Start(notifyChannel chan struct{}, requirements map[string]core.Module) {
	psqlCredits := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.config.host,
		p.config.port,
		p.config.username,
		p.config.password,
		p.config.database,
	)

	db, err := sql.Open("postgres", psqlCredits)
	if err != nil {
		log.Printf("cannot open connection to database: %v", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Printf("cannot ping database: %v", err)
		return
	}

	log.Println("Postgres: successful connection")

	p.db = db

	notifyChannel <- struct{}{}
}

func (p *Postgres) Stop() {
	defer p.db.Close()
}

func (p *Postgres) Require() []string {
	return []string{}
}

func (p *Postgres) Signature() string {
	return "storage"
}
