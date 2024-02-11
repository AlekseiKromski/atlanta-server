package tcp_consumer

type ServerConfig struct {
	Address string
	BufSize int
}

func NewServerConfig(address string, bufSize int) *ServerConfig {
	return &ServerConfig{
		Address: address,
		BufSize: bufSize,
	}
}
