package main

import (
	"context"
	"log"
	"net"
	"os"
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

	for {
		buffer := make([]byte, packetSizeUDP)
		nbytes, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			<-ctx.Done()
			return err
		}

		go func(buffer []byte, remoteAddr net.Addr, nbytes int) {
			log.Printf("packet-received: bytes=%+v from=%+v buffer=%+v", nbytes, remoteAddr, string(buffer))
		}(buffer, addr, nbytes)
	}

	return nil
}
