package tcp_consumer

type ServerConfig struct {
	address string
	bufSize int
}

func NewServerConfig(address string, bufSize int) *ServerConfig {
	return &ServerConfig{
		address: address,
		bufSize: bufSize,
	}
}
