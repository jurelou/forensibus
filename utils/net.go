package utils

import (
	"net"
	"strconv"
	"time"
)

func FindOpenTcpPort() int {
	for port := 50142; port <= 60142; port++ {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", strconv.Itoa(port)), time.Second)
		if err != nil {
			return port
		}
		if conn != nil {
			conn.Close()
		}
	}
	return 0
}
