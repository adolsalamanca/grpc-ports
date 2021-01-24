package main

import (
	"context"
	"flag"
	"github.com/adolsalamanca/ports/client/interface"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	_, cancelFunc := context.WithCancel(context.Background())
	go arrangeShutdown(cancelFunc)

	timeout := flag.Uint("timeout", 5, "Client timeout in seconds to receive data from ports")
	port := flag.Uint("port", 3001, "Client port to receive data from ports")
	host := flag.String("host", "0.0.0.0", "Client host to receive data from ports")
	flag.Parse()

	s := _interface.Server{
		Timeout: time.Duration(*timeout) * time.Second,
		Host:    *host,
		Port:    *port,
	}
	if err := s.Start(context.Background()); err != nil {
		log.Fatalf("cannot start server due to %s", err)
	}

}

func arrangeShutdown(cancelFunc context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan

	log.Printf("Shutting down ports client app...")
	cancelFunc()
	os.Exit(1)
}
