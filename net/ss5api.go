package net

import (
	"io"
	"net"
)

type DataHandlerFunc func(request []byte) ([]byte, error)
type MessageHandlerFunc func(request *Socks5MessageRequest) (*Socks5MessageResponse, error)

type ServerInfo interface {
	Version() byte
	Methods() []byte
}

// BaseServer socks5 server api
type BaseServer interface {
	ServerInfo
	Start() error
	Close() error
	Connection() (SSConn, error)
}

type NoAuthServer interface {
	BaseServer
}

type GssApiServer interface {
	BaseServer
}

// usage username and password authentication
type UserAuthServer interface {
	BaseServer
	SetUserAuth(username string, password string)
}

type SSConn interface {
	io.Closer
	RawConn() net.Conn
	SendResponse(rep *Socks5MessageResponse) error
	ReadRequest() (*Socks5MessageRequest, error)
	ReadMessage() ([]byte, error)
	RegisterConnectHandler(mhf MessageHandlerFunc, dhf DataHandlerFunc)
	Handler() error
}
