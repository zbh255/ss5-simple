package net

import (
	"io"
	"net"
	"reflect"
	"unsafe"
)

const (
	BUFFER_SIZE = 512
)

func newSSConn(conn net.Conn) SSConn {
	return &ssConn{
		rawConn: conn,
		mhf: nil,
		dhf: nil,
	}
}

type ssConn struct {
	rawConn net.Conn
	// mhf is call one
	mhfCallFlags bool
	// handle message
	mhf MessageHandlerFunc
	// handle data
	dhf DataHandlerFunc
}


func (s *ssConn) Close() error {
	return s.rawConn.Close()
}

func (s *ssConn) RawConn() net.Conn {
	return s.rawConn
}

func (s *ssConn) RegisterConnectHandler(mhf MessageHandlerFunc, dhf DataHandlerFunc) {
	s.mhf = mhf
	s.dhf = dhf
}


func (s *ssConn) SendResponse(rep *Socks5MessageResponse) error {
	buf := EncodeSocks5MessageResponse(rep)
	nBytes, err := s.rawConn.Write(buf)
	if err != nil {
		return err
	}
	if nBytes != len(buf) {
		return ErrWriteByteNumberNoEqual
	}
	return nil
}

func (s *ssConn) ReadRequest() (*Socks5MessageRequest, error) {
	request, err := DecodeSocks5MessageRequest(s.rawConn)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (s *ssConn) ReadMessage() ([]byte, error) {
	readN := 0
	buf := make([]byte,BUFFER_SIZE)
	n,err := s.rawConn.Read(buf)
	if err != nil {
		return nil, err
	}
	for n == BUFFER_SIZE {
		readN += n
		buf = append(buf,make([]byte,BUFFER_SIZE)...)
		n, err = s.rawConn.Read(buf[readN : readN+BUFFER_SIZE])
		if err == io.EOF {
			n = 0
			break
		}
		if err != nil {
			return nil, err
		}
	}
	readN += n

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	sh.Len = readN
	return *(*[]byte)(unsafe.Pointer(sh)),nil
}

func (s *ssConn) Handler() error {
	if !s.mhfCallFlags {
		err := s.callMhf(s.mhf)
		if err != nil {
			return err
		}
	}
	return s.callDfh(s.dhf)
}

func (s *ssConn) callMhf(handlerFunc MessageHandlerFunc) error {
	request,err :=s.ReadRequest()
	if err != nil {
		return err
	}
	response,err := handlerFunc(request)
	if err != nil {
		return err
	}
	return s.SendResponse(response)
}

func (s *ssConn) callDfh(handlerFunc DataHandlerFunc) error {
	request,err := s.ReadMessage()
	if err != nil {
		return err
	}
	response,err := handlerFunc(request)
	if err != nil {
		return err
	}
	n,err := s.rawConn.Write(response)
	if err != nil {
		return err
	}
	if n != len(response) {
		return ErrWriteByteNumberNoEqual
	}
	return nil
}

