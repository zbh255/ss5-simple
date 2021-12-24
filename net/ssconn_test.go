package net

import (
	"bytes"
	"math/rand"
	"net"
	"sync"
	"testing"
	"time"
)

func randomByte(n int) []byte {
	buf := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		buf[i] = byte(rand.Intn(255))
	}
	return buf
}

func TestSsConnReadMessage(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	message := randomByte(BUFFER_SIZE * 4)
	go server(t, &wg, message)
	go client(t, &wg, message)
	wg.Wait()
}

func server(t *testing.T, wg *sync.WaitGroup, msg []byte) {
	defer wg.Done()
	listener, err := net.Listen("tcp", "localhost:3456")
	defer listener.Close()
	if err != nil {
		t.Error(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		t.Error(err)
	}
	ssc := newSSConn(conn)
	defer ssc.Close()
	readMessage, err := ssc.ReadMessage()
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(readMessage, msg) {
		t.Error("test message data is not equal client send message")
	}
}

func client(t *testing.T, wg *sync.WaitGroup, msg []byte) {
	defer wg.Done()
	conn, err := net.Dial("tcp", "localhost:3456")
	defer conn.Close()
	if err != nil {
		t.Error(err)
	}
	_, err = conn.Write(msg)
	if err != nil {
		t.Error(err)
	}
}
