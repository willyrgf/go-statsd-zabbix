package main

import "net"

// State is a model of all current state of app
type State struct {
}

// StatsDConfig is a model with all configurations about the statsd
type StatsDServer struct {
	Hostname      string
	Address       string
	Protocol      string
	DefaultPrefix string
	MaxPacketSize int
}

// Datagram is the UDP datagram packet received from the PacketConn
type Datagram struct {
	NumberOfBytes int
	RemoteAddr    net.Addr
	Buffer        []byte
}
