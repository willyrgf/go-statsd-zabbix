package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

// cancelOnInterrupt calls f when os.Interrupt or SIGTERM is received
func cancelOnInterrupt(ctx context.Context, cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		select {
		case <-ctx.Done():
			log.Printf("cancelOnInterrupt() <-ctx.Done(): ctx: %+v c: %+v", ctx, c)
		case <-c:
			log.Printf("cancelOnInterrupt() <-c: ctx: %+v c: %+v", ctx, c)
			cancel()
		}
	}()
}

func readConfig() (*viper.Viper, error) {
	var err error
	v := viper.New()

	v.AutomaticEnv()

	return v, err
}

// run all application
func run() error {
	viper, err := readConfig()
	if err != nil {
		return err
	}

	statsdConfig, err := NewStatsDConfig(viper)
	if err != nil {
		return err
	}

	statsd := NewStatsDServer(statsdConfig)

	// make the context and control then
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	cancelOnInterrupt(ctx, cancelFunc)

	if err := statsd.Run(ctx); err != nil && err != context.Canceled {
		return fmt.Errorf("statsd.Run() error: %+v", err)
	}

	return nil
}

func main() {
	log.Printf("start go-statsd-zabbix")
	if err := run(); err != nil {
		log.Fatalf("main() error run(): %+v\n", err)
	}
}
