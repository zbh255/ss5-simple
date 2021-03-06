package net

const (
	SOCKS5_VERSION byte = 0x05
	SOCKS4_VERSION byte = 0x04
)

const (
	SOCKS5_METHOD_NOAUTH              byte = 0x00
	SOCKS5_METHOD_CSSAPI              byte = 0x01
	SOCKS5_METHOD_USERNAMEANDPASSWORD byte = 0x02
	SOCKS5_METHOD_NOACCEPTABLEMETHODS byte = 0xff
)

const (
	SOCKS5_CMD_CONNECT byte = iota
	SOCKS5_CMD_BIND
	SOCKS5_CMD_UDP_ASSSOCIATE
)

const (
	SOCKS5_ATYP_IPV4   byte = 0x01
	SOCKS5_ATYP_DOMAIN byte = 0x03
	SOCKS5_ATYP_IPV6   byte = 0x04
)

const (
	SOCKS5_REPLY_SUCCESS byte = iota
	SOCKS5_REPLY_FAILURE
	SOCKS5_REPLY_NOT_RULESET
	SOCKS5_REPLY_NETWORK_UNREACHABLE
	SOCKS5_REPLY_HOST_UNREACHABLE
	SOCKS5_REPLY_CONNECTION_REFUSED
	SOCKS5_REPLY_TTL_EXPIRED
	SOCKS5_REPLY_CMD_NOTSUPPORT
	SOCKS5_REPLY_ADDRESSTYPE_NOTSUPPORT
	SOCKS5_REPLY_UNASSIGNED
)
