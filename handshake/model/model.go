package model

type Config struct {
	Host string `env:"host"`
	Port int    `env:"port"`
}

type HandshakeRequest struct {
	Username        string
	HandshakeStatus string
	HandshakeAt     string
}

type HandshakeReply struct {
	Message string
}
