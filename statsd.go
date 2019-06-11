package main

import (
	"net"
	"os"
)

const (
	maxPacketSize = 1400
)

// StatsDSocketFactory is a layer over net.ListenPacket() to allow
// implementations
type StatsDSocketFactory func() (net.PacketConn, error)

func socketFactory(protocol, addr string) StatsDSocketFactory {
	conn, err := net.ListenPacket(protocol, addr)
	return func() (net.PacketConn, error) {
		return conn, err
	}
}

// NewStatsDServer get struct encapsulate all of parameters
// for start the statsd server
func NewStatsDServer() *StatsDServer {
	h, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return &StatsDServer{
		Hostname:      h,
		Address:       "127.0.0.1:8125",
		Protocol:      "udp",
		DefaultPrefix: "",
		MaxPacketSize: maxPacketSize,
	}
}
