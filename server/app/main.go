package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	_, cancelFunc := context.WithCancel(context.Background())
	go arrangeShutdown(cancelFunc)

	go func() {
		for {
			t := time.NewTimer(2 * time.Second)
			<-t.C
			log.Printf("Server app running...")
		}

	}()

	waitChan := make(chan bool,1)
	<- waitChan
}

func arrangeShutdown(cancelFunc context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-signalChan

	log.Printf("Shutting down ports server app...")
	cancelFunc()
	os.Exit(1)
}
