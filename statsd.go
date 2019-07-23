package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/spf13/viper"
)

const (
	packetSizeUDP = 0xffff
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

// NewStatsDConfig start the StatsDConfig of the app with
// all configurations about the statsd server
func NewStatsDConfig(viper *viper.Viper) (StatsDConfig, error) {
	var err error
	config := StatsDConfig{}

	config.Hostname = viper.GetString("HOSTNAME")
	if len(config.Hostname) < 1 {
		config.Hostname, err = os.Hostname()
	}

	config.StatsDPrefix = viper.GetString("STATSD_PREFIX")
	config.StorageType, err = NewStorageType(viper.GetString("STORAGE_TYPE"))
	config.StorageURL = viper.GetString("STORAGE_URL")

	return config, err
}

// NewStatsDServer get struct encapsulate all of parameters
// for start the statsd server
func NewStatsDServer(config StatsDConfig) *StatsDServer {
	h, err := os.Hostname()
	if err != nil {
		log.Printf("NewStatsDServer(): GetHostname() error: %v", err)
		os.Exit(1)
	}

	// initialization cache
	cache, err := NewStorage(Memory, "")
	if err != nil {
		log.Printf("NewStatsDServer(): NewStorage() error: %v", err)
		os.Exit(1)
	}

	// initialization of storage
	storage, err := NewStorage(config.StorageType, config.StorageURL)
	if err != nil {
		log.Printf("NewStatsDServer(): NewStorage() error: %v", err)
		os.Exit(1)
	}

	return &StatsDServer{
		Hostname:      h,
		Address:       "127.0.0.1:8125",
		Protocol:      "udp",
		DefaultPrefix: "",
		Storage:       storage,
		Cache:         cache,
		Config:        config,
	}
}

// Run runs the server until context signals done
func (s *StatsDServer) Run(ctx context.Context) error {
	return s.RunWithSocket(ctx, socketFactory(s.Protocol, s.Address))
}

// RunWithSocket runs the server until context signals done
// listering socket is created using socket
func (s *StatsDServer) RunWithSocket(ctx context.Context, socket StatsDSocketFactory) error {
	conn, err := socket()
	if err != nil {
		<-ctx.Done()
		return err
	}
	defer conn.Close()

	doneReceiver := make(chan error, 1)
	doneRunMetrics := make(chan error, 1)
	receivedDatagram := make(chan Datagram)
	go ReceiverDatagram(ctx, conn, doneReceiver, receivedDatagram)
	go RunMetrics(ctx, doneRunMetrics, receivedDatagram, s)

	select {
	case <-ctx.Done():
	case <-doneReceiver:
	}

	return nil
}
