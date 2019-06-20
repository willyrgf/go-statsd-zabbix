package main

import (
	"os"

	"github.com/spf13/viper"
)

// NewState start the state of the app with
// all configurations about the statsd server
func NewState(viper *viper.Viper) (State, error) {
	var err error
	state := State{}

	h, err := os.Hostname()

	state.Hostname = h
	state.StatsDPrefix = viper.GetString("STATSD_PREFIX")
	state.StorageType = viper.GetString("STORAGE_TYPE")

	return state, err
}
