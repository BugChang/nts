package model

import "net"

// ClientInfo 客户端信息
type ClientInfo struct {
	ID   int
	Addr string
	Conn net.Conn
}
