package listen

import (
	"bufio"
	"fmt"
	"net"

	"github.com/bugchang/nts/model"
)

var clients = make(map[int]model.ClientInfo)

// ListenClient 监听客户端连接
func ListenClient() {
	listen, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	fmt.Println("服务启动成功，监听地址：127.0.0.1:5000")
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		remoteAddr := conn.RemoteAddr().String()
		fmt.Printf("客户端:%s连接成功\r\n", remoteAddr)
		clients[1] = model.ClientInfo{
			ID:   1,
			Addr: remoteAddr,
			Conn: conn,
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}

// 处理函数
func process(conn net.Conn) {
	// defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
	}
}
