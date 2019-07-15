package main

import (
	"net"
	"time"
)

// StatsDConfig is a model of all configs of app
type StatsDConfig struct {
	StorageType  StorageType
	StorageURL   string
	Hostname     string
	StatsDPrefix string
}

// StatsDServer is a model with all configurations about the statsd
type StatsDServer struct {
	Hostname      string
	Address       string
	Protocol      string
	DefaultPrefix string
	MaxPacketSize int
	Config        StatsDConfig
}

// Datagram is the UDP datagram packet received from the PacketConn
type Datagram struct {
	NumberOfBytes int
	RemoteAddr    net.Addr
	Buffer        []byte
}

// StatsDMetric is a pure representation for a  statsd metric parsed from UDP datagram
type StatsDMetric struct {
	Name    string  `json:"name"`
	NameRaw string  `json:"name_raw"`
	Value   float64 `json:"value"`
	Type    string  `json:"type"`
}

// Metric is a representation of all data collected with in each UDP datagram
type Metric struct {
	Hostname   string       `json:"hostname"`
	SourceIP   string       `json:"source_ip"`
	SourcePort string       `json:"source_port"`
	Timestamp  time.Time    `json:"timestamp"`
	Prefix     string       `json:"prefix"`
	Stats      StatsDMetric `json:"stats"`
}
