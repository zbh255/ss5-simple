package net

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestHandShake(t *testing.T) {
	requestBytes := []byte{
		SOCKS5_VERSION, // socks version
		0x01,           // n methods
		// method table
		0x01,
	}
	reader := bytes.NewReader(requestBytes)
	request, err := DecodeSocks5HandshakeRequest(reader)
	if err != nil {
		t.Error(err)
	}
	t.Log(request)
	// handshake response
	response := new(Socks5HandshakeResponse)
	response.Version = request.Version
	response.Method = SOCKS5_METHOD_NOAUTH
	responseBytes := EncodeSocks5HandshakeResponse(response)
	t.Log(responseBytes)
}

func TestConnect(t *testing.T) {
	bigEndianPort := make([]byte, 2)
	domain := []byte("cdn.blog.baidu.com")
	binary.BigEndian.PutUint16(bigEndianPort, 8080)
	nRequestBytes := [][]byte{
		{
			SOCKS5_VERSION,     // socks version,
			SOCKS5_CMD_CONNECT, // command type
			0x00,               // rsv
			SOCKS5_ATYP_IPV4,   // address type
			192,                // ipv4 1 byte
			168,                // ipv4 2 byte
			10,                 // ipv4 3 byte
			9,                  // ipv4 4 byte
			bigEndianPort[0],   // big endian port
			bigEndianPort[1],   // big endian port
		},
		{
			SOCKS5_VERSION,     // socks version,
			SOCKS5_CMD_CONNECT, // command type
			0x00,               // rsv
			SOCKS5_ATYP_DOMAIN, // address type
			byte(len(domain)),  // domain length max 256
		},
		{
			SOCKS5_VERSION,     // socks version,
			SOCKS5_CMD_CONNECT, // command type
			0x00,               // rsv
			SOCKS5_ATYP_IPV6,   // address type
			0xff,               // ipv6 1 byte
			0xff,               // ipv6 2 byte
			0xff,               // ipv6 3 byte
			0xff,               // ipv6 4 byte
			0xff,               // ipv6 5 byte
			0xff,               // ipv6 6 byte
			0xff,               // ipv6 7 byte
			0xff,               // ipv6 8 byte
			0xff,               // ipv6 9 byte
			0xff,               // ipv6 10 byte
			0xff,               // ipv6 11 byte
			0xff,               // ipv6 12 byte
			0xff,               // ipv6 13 byte
			0xff,               // ipv6 14 byte
			0xff,               // ipv6 15 byte
			0xff,               // ipv6 16 byte
			bigEndianPort[0],   // big endian port
			bigEndianPort[1],   // big endian port
		},
	}
	// domain element join
	nRequestBytes[1] = append(nRequestBytes[1], domain...)
	nRequestBytes[1] = append(nRequestBytes[1], bigEndianPort...)
	nRequest := make([]*Socks5MessageRequest, 0, 3)
	for _, v := range nRequestBytes {
		reader := bytes.NewReader(v)
		request, err := DecodeSocks5MessageRequest(reader)
		if err != nil {
			t.Error(err)
		}
		nRequest = append(nRequest, request)
		t.Log(request)
	}
	// response testing
	replyMapping := map[int]byte{
		0: SOCKS5_REPLY_SUCCESS,
		1: SOCKS5_REPLY_TTL_EXPIRED,
		2: SOCKS5_REPLY_UNASSIGNED,
	}
	var nResponse []*Socks5MessageResponse
	for k, v := range nRequest {
		rep := new(Socks5MessageResponse)
		rep.Version = v.Version
		rep.Reply = replyMapping[k]
		rep.Reserved = 0x00
		rep.AddrType = v.AddrType
		rep.Adders = v.Adders
		rep.Port = v.Port

		nResponse = append(nResponse, rep)
	}

	for _, v := range nResponse {
		responseBytes := EncodeSocks5MessageResponse(v)
		t.Log(responseBytes)
	}
}
