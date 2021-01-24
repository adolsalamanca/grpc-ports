package main

import (
	"context"
	"flag"
	"github.com/adolsalamanca/ports/server/infrastructure/persistence"
	"github.com/adolsalamanca/ports/server/interface/server"
	"google.golang.org/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_, cancelFunc := context.WithCancel(context.Background())
	go arrangeShutdown(cancelFunc)

	log.Printf("Server app starting...")
	port := flag.Uint("port", 7777, "Port of gRpc server")
	flag.Parse()

	grpcServer := grpc.NewServer()
	portsRepository := persistence.NewPortsMemoryRepository()
	s := server.NewServer(grpcServer, portsRepository)

	s.Serve(*port)
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
