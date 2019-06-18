package main

import (
	"context"
	"log"
	"net"
)

func NewDatagram(buffer []byte, addr net.Addr, nbytes int) Datagram {
	return Datagram{
		NumberOfBytes: nbytes,
		RemoteAddr:    addr,
		Buffer:        buffer,
	}
}

func ReceiverDatagram(ctx context.Context, conn net.PacketConn, done chan<- error, receivedDatagram chan Datagram) {
	for {
		buffer := make([]byte, packetSizeUDP)
		nbytes, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			<-ctx.Done()
			done <- err
		}

		d := NewDatagram(buffer, addr, nbytes)
		receivedDatagram <- d
	}

}

func RunMetrics(ctx context.Context, done chan<- error, receivedDatagram <-chan Datagram) {
	//log.Printf("packet-received: bytes=%+v from=%+v buffer=%+v", nbytes, remoteAddr, string(buffer))

	for {
		select {
		case d := <-receivedDatagram:
			log.Printf("chan receivedDatagram: %+v", d)
		}
	}
}
