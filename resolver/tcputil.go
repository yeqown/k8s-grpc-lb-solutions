package resolver

import (
	"fmt"
	"net"
)

func Sniff(address string) string {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("error=%v", err)
		return ""
	}
	defer conn.Close()

	tcpconn := conn.(*net.TCPConn)
	return tcpconn.RemoteAddr().String()
}
