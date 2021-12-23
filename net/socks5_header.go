package net

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type Socks5HandshakeRequest struct {
	Version  byte
	NMethods byte
	Methods  [256]byte
}

func DecodeSocks5HandshakeRequest(reader io.Reader) (*Socks5HandshakeRequest, error) {
	sock := new(Socks5HandshakeRequest)
	buf := make([]byte, 2)
	_, err := reader.Read(buf)
	if err != nil {
		return nil, err
	}
	sock.Version = buf[0]
	sock.NMethods = buf[1]
	buf = make([]byte, sock.NMethods)
	_, err = reader.Read(buf)
	if err != nil {
		return nil, err
	}
	copy(sock.Methods[:sock.NMethods], buf)
	return sock, nil
}

type Socks5HandshakeResponse struct {
	Version byte
	Method  byte
}

func EncodeSocks5HandshakeResponse(response *Socks5HandshakeResponse) []byte {
	buf := make([]byte, 2)
	buf[0] = response.Version
	buf[1] = response.Method
	return buf
}

type Socks5MessageRequest struct {
	Version  byte
	Command  byte
	Reserved byte
	AddrType byte
	Adders   []byte
	Port     uint16
}

func DecodeSocks5MessageRequest(reader io.Reader) (*Socks5MessageRequest, error) {
	buf := make([]byte, 4)
	_, err := reader.Read(buf)
	if err != nil {
		return nil, err
	}
	sock := new(Socks5MessageRequest)
	sock.Version = buf[0]
	sock.Command = buf[1]
	sock.Reserved = 0x00
	sock.AddrType = buf[3]

	switch sock.AddrType {
	case SOCKS5_ATYP_IPV4:
		buf = make([]byte, 4)
	case SOCKS5_ATYP_IPV6:
		buf = make([]byte, 16)
	case SOCKS5_ATYP_DOMAIN:
		buffer := make([]byte, 1)
		_, err := reader.Read(buffer)
		if err != nil {
			return nil, err
		}
		domainLength := buffer[0]
		buf = make([]byte, domainLength)
	default:
		return nil, errors.New(fmt.Sprintf("socks5 address type not supported : %d", sock.AddrType))
	}
	_, err = reader.Read(buf)
	if err != nil {
		return nil, err
	}
	sock.Adders = buf
	// read port
	buf = make([]byte, 2)
	_, err = reader.Read(buf)
	if err != nil {
		return nil, err
	}
	sock.Port = binary.BigEndian.Uint16(buf)
	return sock, nil
}

type Socks5MessageResponse struct {
	Version  byte
	Reply    byte
	Reserved byte
	AddrType byte
	Adders   []byte
	Port     uint16
}

func EncodeSocks5MessageResponse(response *Socks5MessageResponse) []byte {
	buf := make([]byte, 0, 16)
	buf = append(buf, response.Version)
	buf = append(buf, response.Reply)
	buf = append(buf, response.Reserved)
	buf = append(buf, response.AddrType)
	// domain
	if response.AddrType == SOCKS5_ATYP_DOMAIN {
		buf = append(buf, byte(len(response.Adders)))
	}
	buf = append(buf, response.Adders...)
	buf = append(buf, make([]byte, 2)...)
	binary.BigEndian.PutUint16(buf[len(buf)-2:], response.Port)
	return buf
}
