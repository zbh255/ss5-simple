package main

import (
	"github.com/zbh255/ss5-simple/handler"
	snet "github.com/zbh255/ss5-simple/net"
	"io/ioutil"
	"log"
	"net"
)

func SimpleServer() {
	listener, err := net.Listen("tcp", "localhost:1080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	socks5Server := snet.NewNoAuthServer(listener)
	for {
		ssc, err := socks5Server.Connection()
		if err != nil {
			log.Print(err.Error())
			continue
		}
		go func() {
			err := simpleHandlerConnection(ssc)
			if err != nil {
				log.Printf("[Error]: %v", err)
			}
		}()
	}
}

func simpleHandlerConnection(conn snet.SSConn) error {
	rep,err := ioutil.ReadFile("./hello.html")
	if err != nil {
		return err
	}
	defer conn.Close()
	conn.RegisterConnectHandler(handler.Comment, func(request []byte) ([]byte, error) {
		return rep, nil
	})
	return conn.Handler()
}
