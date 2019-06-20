package main

import (
	"net"
	"time"
)

// State is a model of all current state of app
type State struct {
	StorageType  string
	Hostname     string
	StatsDPrefix string
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

// StatsDMetric is a pure representation for a  statsd metric parsed from UDP datagram
type StatsDMetric struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

// Metric is a representation of all data collected with in each UDP datagram
type Metric struct {
	SourceIP   string       `json:"source_ip"`
	SourcePort string       `json:"source_port"`
	Timestamp  time.Time    `json:"timestamp"`
	Prefix     string       `json:"prefix"`
	Stats      StatsDMetric `json:"stats"`
}
