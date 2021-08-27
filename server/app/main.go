package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adolsalamanca/grpc-ports/server/infrastructure/persistence"
	"github.com/adolsalamanca/grpc-ports/server/interface/server"
	"google.golang.org/grpc"
)

func main() {
	_, cancelFunc := context.WithCancel(context.Background())
	go arrangeShutdown(cancelFunc)

	log.Printf("Server app starting...")
	port := flag.Uint("port", 3002, "Port of gRpc server")
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
