package util

import (
	"fmt"
	"net"
)

func GenRandomPort() int {
	// 返回 *TCPAddr
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	fmt.Println(addr)
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	port := listener.Addr().(*net.TCPAddr).Port
	return port
}
