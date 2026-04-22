package main

import (
	"fmt"
	"net"
	"time"
)

type Scanner struct {
	Timeout time.Duration
	Rate    time.Duration
}

func (s *Scanner) ScanPort(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", address, s.Timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
