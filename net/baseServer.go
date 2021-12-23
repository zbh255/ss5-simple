package net

import (
	"net"
)

type baseServer struct {
	version     byte
	methodTable []byte
	listener    net.Listener
}

func (b *baseServer) Version() byte {
	return b.version
}

func (b *baseServer) Methods() []byte {
	return b.methodTable
}

func (b *baseServer) Start() error {
	return nil
}

func (b *baseServer) Close() error {
	return b.listener.Close()
}

func (b *baseServer) Connection() (SSConn, error) {
	conn, err := b.listener.Accept()
	if err != nil {
		return nil, err
	}
	// handshake
	handshakeRequest, err := DecodeSocks5HandshakeRequest(conn)
	if err != nil {
		return nil, err
	}
	handshakeResponse := new(Socks5HandshakeResponse)
	handshakeResponse.Version = handshakeRequest.Version
	handshakeResponse.Method = SOCKS5_METHOD_NOAUTH
	// ssConn
	ssconn := newSSConn(conn)
	buf := EncodeSocks5HandshakeResponse(handshakeResponse)
	n, err := ssconn.RawConn().Write(buf)
	if err != nil {
		return nil, err
	}
	if n != len(buf) {
		return nil, ErrWriteByteNumberNoEqual
	}
	return ssconn, nil
}
