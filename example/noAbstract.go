package main

import (
	"errors"
	snet "github.com/zbh255/ss5-simple/net"
	"io/ioutil"
	"log"
	"net"
)

func NoAbstractServer() {
	listener,err := net.Listen("tcp","127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn,err := listener.Accept()
		if err != nil {
			log.Printf("[Error] : %s\n", err.Error())
		}
		go func() {
			err := handlerConnection(conn)
			if err != nil {
				log.Printf("[Error] : %s\n", err.Error())
			}
		}()
	}
}

func handlerConnection(conn net.Conn) error {
	// custom defined response message
	defer conn.Close()
	rep, err := ioutil.ReadFile("./hello.html")
	if err != nil {
		panic(err)
	}
	// from conn read handshake request
	hRequest, err := snet.DecodeSocks5HandshakeRequest(conn)
	if err != nil {
		return err
	}
	// create handshake response
	hResponse := new(snet.Socks5HandshakeResponse)
	hResponse.Version = hRequest.Version
	hResponse.Method = snet.SOCKS5_METHOD_NOAUTH
	hResponseBytes := snet.EncodeSocks5HandshakeResponse(hResponse)
	n,err := conn.Write(hResponseBytes)
	if err != nil {
		return err
	}
	if n != len(hResponseBytes) {
		return errors.New("write bytes no equal")
	}
	// read client message request
	mRequest,err := snet.DecodeSocks5MessageRequest(conn)
	if err != nil {
		return err
	}
	// create response message
	mResponse := new(snet.Socks5MessageResponse)
	mResponse.Version = mRequest.Version
	mResponse.Reply = snet.SOCKS5_REPLY_SUCCESS
	mResponse.Reserved = mRequest.Reserved
	mResponse.AddrType = mRequest.AddrType
	mResponse.Adders = mRequest.Adders
	mResponse.Port = mRequest.Port
	mResponseBytes := snet.EncodeSocks5MessageResponse(mResponse)
	n,err = conn.Write(mResponseBytes)
	if err != nil {
		return err
	}
	if n != len(mResponseBytes) {
		return errors.New("write bytes no equal")
	}
	// server ok
	// read bytes message
	buffer := make([]byte,4096)
	_, err = conn.Read(buffer)
	if err != nil {
		return err
	}
	_, err = conn.Write(rep)
	return err
}
