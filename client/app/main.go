package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adolsalamanca/grpc-ports/client/infrastructure/persistence"
	_interface "github.com/adolsalamanca/grpc-ports/client/interface"
)

func main() {
	_, cancelFunc := context.WithCancel(context.Background())
	go arrangeShutdown(cancelFunc)

	timeout := flag.Uint("timeout", 5, "Client timeout in seconds to receive data from ports")
	port := flag.Uint("port", 3001, "Client port to receive data from ports")
	grpcServerHost := flag.String("grpc_server_host", "0.0.0.0", "host of gRpc server")
	grpcServerPort := flag.Uint("grpc_server_port", 3002, "Port of gRpc server")
	host := flag.String("host", "0.0.0.0", "Client host to receive data from ports")
	flag.Parse()

	s := _interface.Server{
		Timeout: time.Duration(*timeout) * time.Second,
		Host:    *host,
		Port:    *port,
	}

	repository := persistence.NewPortgRpcPersistence(*grpcServerHost, *grpcServerPort)
	if err := s.Start(context.Background(), repository); err != nil {
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
