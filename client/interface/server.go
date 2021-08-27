package _interface

import (
	"context"
	"github.com/adolsalamanca/grpc-ports/client/domain/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net"
	"strconv"
	"time"
)

type Server struct {
	Host    string
	Port    uint
	Timeout time.Duration
}

// Start creates an instance of Echo, high performance, extensible, minimalist Go web framework.
func (s Server) Start(ctx context.Context, repository repository.PortRepository) error {
	e := s.echo(repository)

	go func() {
		<-ctx.Done()

		if err := e.Shutdown(ctx); err != nil {
			log.Error(err)
		}
	}()

	return e.Start(s.Addr())
}

func (s Server) echo(repository repository.PortRepository) *echo.Echo {
	log.Printf("Http listener running on %s", s.Addr())
	e := echo.New()
	e.Use(middleware.Recover())

	p := NewPortHandler(repository)

	e.GET("/grpc-ports", p.GetPorts)
	e.POST("/grpc-ports", p.StorePorts)

	return e
}

// Addr of HTTP server (Host:Port).
func (s Server) Addr() string {
	return net.JoinHostPort(s.Host, strconv.Itoa(int(s.Port)))
}
