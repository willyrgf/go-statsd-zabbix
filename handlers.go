package main

import (
	"bytes"
	"context"
	"errors"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

// NewDatagram returns the Datagram struct completed
func NewDatagram(buffer []byte, addr net.Addr, nbytes int) Datagram {
	return Datagram{
		NumberOfBytes: nbytes,
		RemoteAddr:    addr,
		Buffer:        buffer,
	}
}

// ReceiverDatagram receive the UDP datagram and return
// to the channel the Datagram received
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

// splitNameMetric split the raw name of metric to human name to this metric
func splitNameMetric(n string) string {
	var nameSplitted []string
	var suffix string

	nameSplitted = strings.Split(n, ".")
	if len(nameSplitted) > 0 {
		suffix = "." + nameSplitted[len(nameSplitted)-1]
	}

	nameSplitted = strings.Split(n, ";")
	if len(nameSplitted) >= 2 {
		return nameSplitted[0] + suffix
	}

	nameSplitted = strings.Split(n, "@")
	if len(nameSplitted) >= 2 {
		return nameSplitted[0] + suffix
	}

	return n
}

// handleDatagram handle the msg from udp datagram packet
func handleDatagram(d Datagram) (nameRaw string, name string, value float64, typeOf string, err error) {
	var msg []byte

	if len(d.Buffer) == 0 {
		err = errors.New("the length of the buffer of datagram is zero")
		return
	}

	idx := bytes.IndexByte(d.Buffer, '\n')
	// protocol does not require line to end in \n
	if idx == -1 { // \n not found
		msg = d.Buffer[:d.NumberOfBytes]
	}

	msg = d.Buffer[:idx]

	sMsg := string(msg)
	splittedMsg := strings.Split(sMsg, ":")
	splittedValueType := strings.Split(splittedMsg[1], "|")

	nameRaw = splittedMsg[0]
	name = splitNameMetric(nameRaw)
	typeOf = splittedValueType[1]
	value, err = strconv.ParseFloat(splittedValueType[0], 64)

	return
}

// ParseStatsDMetric parse the Datagram received to a StatsDMetric
func (d Datagram) ParseStatsDMetric() (StatsDMetric, error) {
	stats := StatsDMetric{}
	var err error

	stats.NameRaw, stats.Name, stats.Value, stats.Type, err = handleDatagram(d)

	log.Printf("ParseStatsDMetric: stats: %+v\n", stats)

	return stats, err
}

// ParseMetric parse the Datagram received to a Metric
// with all fields completed and ready for storage
func (d Datagram) ParseMetric(statsd *StatsDServer) (Metric, error) {
	m := Metric{}
	var err error

	m.Hostname = statsd.Config.Hostname
	m.SourceIP, m.SourcePort, err = net.SplitHostPort(d.RemoteAddr.String())
	m.Timestamp = time.Now()
	m.Prefix = ""

	m.Stats, err = d.ParseStatsDMetric()

	return m, err
}

// Save store the metric in the storage previous configured
func (metric Metric) Save(statsd *StatsDServer) {
	storage, err := NewStorage(statsd.Config.StorageType, statsd.Config.StorageURL)
	if err != nil {
		log.Printf("metric.Save(): err=%+v\n", err)
		return
	}

	if err := storage.SaveMetric(metric); err != nil {
		log.Printf("metric.Save(): err=%+v\n", err)
		return
	}
}

// RunMetrics receive all Datagram by a channel and run all operations
// for proccess and save/storage this metric
func RunMetrics(ctx context.Context, done chan<- error, receivedDatagram <-chan Datagram, statsd *StatsDServer) {
	for {
		select {
		case <-ctx.Done():
			return
		case d := <-receivedDatagram:
			metric, err := d.ParseMetric(statsd)
			if err != nil {
				done <- err
			}
			go metric.Save(statsd)
		}
	}
}
