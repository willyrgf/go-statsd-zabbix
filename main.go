package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// cancelOnInterrupt calls f when os.Interrupt or SIGTERM is received
func cancelOnInterrupt(ctx context.Context, f context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-ctx.Done():
		case <-c:
			f()
		}
	}()
}

// run all application
func run() error {
	statsd := NewStatsDServer()

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
