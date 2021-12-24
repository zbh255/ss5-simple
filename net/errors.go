package net

import "errors"

var (
	ErrWriteByteNumberNoEqual = errors.New("write byte number no equal")
	ErrNotSupportCmd          = errors.New("client request command not support")
)
