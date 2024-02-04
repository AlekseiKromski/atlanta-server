package core

import (
	"log"
	"os"
	"os/signal"
	"time"
)

type Core struct {
	systemChannel chan os.Signal
	notifyChannel chan struct{}
}

func NewCore() *Core {
	return &Core{
		systemChannel: make(chan os.Signal, 1),
		notifyChannel: make(chan struct{}, 1),
	}
}

func (c *Core) Init(modules []Module) {
	// Subscribe to Interrupt syscall
	signal.Notify(c.systemChannel, os.Interrupt)

	c.startModules(modules)

	// Block main thread until stop
	<-c.systemChannel

	c.stopModules(modules)

	// Exit from application
	os.Exit(0)
}

func (c *Core) startModules(modules []Module) {
	startTime := time.Now()
	log.Printf("Core: Start %d modules", len(modules))
	for _, module := range modules {
		go module.Start(c.notifyChannel)

		//Wait until module start
		<-c.notifyChannel
	}
	log.Printf("Core: All modules started in %f seconds", time.Now().Sub(startTime).Seconds())
}

func (c *Core) stopModules(modules []Module) {
	startTime := time.Now()
	log.Printf("Core: will stop %d modules", len(modules))

	for _, module := range modules {
		module.Stop()
	}

	log.Printf("Core: All modules stopped in %f seconds", time.Now().Sub(startTime).Seconds())
}
