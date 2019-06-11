package main

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
