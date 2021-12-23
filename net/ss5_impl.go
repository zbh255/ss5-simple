package net

import "net"

func NewNoAuthServer(listener net.Listener) NoAuthServer {
	return NewBaseServer(listener)
}

//func NewGssApiServer() GssApiServer {}
//
//func NewUserAuthServer() UserAuthServer {}

func NewBaseServer(listener net.Listener) BaseServer {
	return &baseServer{
		version:     SOCKS5_VERSION,
		methodTable: []byte{},
		listener:    listener,
	}
}
