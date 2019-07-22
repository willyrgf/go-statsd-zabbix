package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/adubkov/go-zabbix"
)

const (
	keyOfItem             = `asterisk.statsd.discovery`
	prefixKeyOfMetric     = `asterisk.statsd.metrics`
	minMetricNameCountSep = 2
	separatorOfMetric     = `.`
)

// StorageZabbixSender is the data storage layer using to send data to Zabbix Server
type StorageZabbixSender struct {
	Sender  *zabbix.Sender
	Packet  *zabbix.Packet
	Metrics []*zabbix.Metric
}

// parseTimestampToInt64
func parseTimestampToInt64(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// parseMetricName
func parseMetricToItemValue(metricName string) (valueOfMetric string) {
	return fmt.Sprintf("{\"data\": [{\"{#ITEM}\": \"%s\"}]}", strings.ToLower(metricName))
}

// parseMetricToKey
func parseMetricToKey(metricName string) (key string) {
	var suffix string

	splitted := strings.Split(strings.ToLower(metricName), separatorOfMetric)

	// if th metric name dont have minimun count separators, use the first
	if len(splitted) <= minMetricNameCountSep {
		suffix = separatorOfMetric + splitted[0]
	} else {
		for i, k := range splitted {
			if i >= minMetricNameCountSep {
				break
			}
			suffix = suffix + separatorOfMetric + k
		}
	}

	key = fmt.Sprintf("%s[\"%s\"]", prefixKeyOfMetric+suffix, metricName)
	return
}

func parseMetricValue(metric Metric) (value string) {
	switch metric.Stats.Type {
	case "ms":
		seconds := metric.Stats.Value / 1000
		value = fmt.Sprintf("%f", seconds)
	default:
		value = fmt.Sprintf("%f", metric.Stats.Value)
	}
	return
}

// parseURL parse the url configure to zabbix sender
func parseURL(zabbixURL string) (host, port string, err error) {
	u, err := url.Parse(zabbixURL)
	if err != nil {
		return
	}

	if u.Scheme != "zabbix" {
		err = fmt.Errorf("URL scheme is not Zabbix, is: %s", u.Scheme)
		return
	}

	host, port, err = net.SplitHostPort(u.Host)
	return
}

// NewStorageZabbixSender prepare a Zabbix Sender
func NewStorageZabbixSender(storageURL string) (*StorageZabbixSender, error) {
	var err error
	stg := new(StorageZabbixSender)

	host, p, err := parseURL(storageURL)
	if err != nil {
		return stg, err
	}

	port, err := strconv.Atoi(p)
	if err != nil {
		return stg, err
	}

	stg.Sender = zabbix.NewSender(host, port)

	return stg, err
}

// sendMetrics send the metrics in zabbix storage
func sendMetrics(s *StorageZabbixSender) {
	for _, m := range s.Metrics {
		log.Printf("sendMetrics: m=%+v", m)
	}
	res, err := s.Sender.Send(s.Packet)
	log.Printf("sendMetrics res=%+v, err=%+v", string(res), err)
	return
}

// SaveMetric save the metric on a item in zabbix server
func (s *StorageZabbixSender) SaveMetric(metric Metric) error {
	time := parseTimestampToInt64(metric.Timestamp)
	key := parseMetricToKey(metric.Stats.Name)
	value := parseMetricValue(metric)

	s.Metrics = append(s.Metrics, zabbix.NewMetric(metric.Hostname, key, value, time))

	s.Packet = zabbix.NewPacket(s.Metrics)

	sendMetrics(s)

	return nil
}

// SaveItem save/create the item on zabbix server
func (s *StorageZabbixSender) SaveItem(metric Metric) error {
	time := parseTimestampToInt64(metric.Timestamp)
	value := parseMetricToItemValue(metric.Stats.Name)

	s.Metrics = append(s.Metrics, zabbix.NewMetric(metric.Hostname, keyOfItem, value, time))

	return nil
}

// ItemExists check if a item exists, but it's not implemented to this storage
func (s *StorageZabbixSender) ItemExists(metric Metric) (found bool) {
	return
}
