package main

import (
	"github.com/adubkov/go-zabbix"
)

// StorageZabbixSender is the data storage layer using to send data to Zabbix Server
type StorageZabbixSender struct {
	sender *zabbix.Sender
	packet *zabbix.Packet
	metric *zabbix.Metric
}

// NewStorageZabbixSender prepare a Zabbix Sender
func NewStorageZabbixSender() (*StorageZabbixSender, error) {
	var err error
	stg := new(StorageZabbixSender)

	return stg, err
}

// SaveMetric save the metric on a item in zabbix server
func (s *StorageZabbixSender) SaveMetric(metric Metric) error {
	return nil
}

// SaveItem save/create the item on zabbix server
func (s *StorageZabbixSender) SaveItem(metric Metric) error {
	return nil
}

// ItemExists check if a item exists, but it's not implemented to this storage
func (s *StorageZabbixSender) ItemExists(metric Metric) (found bool) {
	return
}
