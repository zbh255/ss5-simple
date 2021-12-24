package handler

import snet "github.com/zbh255/ss5-simple/net"

func Comment(request *snet.Socks5MessageRequest) (response *snet.Socks5MessageResponse, err error) {
	response = new(snet.Socks5MessageResponse)
	response.Version = request.Version
	response.Reply = snet.SOCKS5_REPLY_SUCCESS
	response.Reserved = request.Reserved
	response.AddrType = request.AddrType
	response.Adders = request.Adders
	response.Port = request.Port
	return
}

// Bind not supported
func Bind(request *snet.Socks5MessageRequest) (response *snet.Socks5MessageResponse, err error) {
	response, err = Comment(request)
	err = snet.ErrNotSupportCmd
	response.Reply = snet.SOCKS5_REPLY_CMD_NOTSUPPORT
	return
}

// UdpAssociate not supported
func UdpAssociate(request *snet.Socks5MessageRequest) (response *snet.Socks5MessageResponse, err error) {
	response, err = Bind(request)
	return
}
