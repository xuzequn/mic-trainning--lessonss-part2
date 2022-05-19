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
	lintener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer lintener.Close()
	port := lintener.Addr().(*net.TCPAddr).Port
	return port
}
