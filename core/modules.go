package core

type Module interface {
	Start(notifyChannel chan struct{}) // Start module
	Stop()                             // Stop module
}
