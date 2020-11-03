package server

import (
	"fmt"
	"net"

	"github.com/bugchang/nts/server/listen"
)

var clients = make(map[string]net.Conn)

// Run 启动服务端
func Run() {
	listen.ListenClient()
}

func distribute(conn net.Conn, recvStr string) {
	var currentAddr = conn.RemoteAddr().String()
	sendStr := fmt.Sprintf("来自客户端【%v】的消息：%s", currentAddr, recvStr)
	for addr, client := range clients {
		if currentAddr != addr {
			client.Write([]byte(sendStr))
		}
	}
}
