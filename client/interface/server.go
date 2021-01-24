package _interface

import (
	"context"
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
func (s Server) Start(ctx context.Context) error {
	e := s.echo()

	go func() {
		<-ctx.Done()

		if err := e.Shutdown(ctx); err != nil {
			log.Error(err)
		}
	}()

	return e.Start(s.Addr())
}

func (s Server) echo() *echo.Echo {
	log.Printf("Http listener running on %s", s.Addr())
	e := echo.New()
	e.Use(middleware.Recover())

	p := PortHandler{}

	e.GET("/ports", p.GetPorts)
	e.POST("/ports", p.StorePorts)

	return e
}

// Addr of HTTP server (Host:Port).
func (s Server) Addr() string {
	return net.JoinHostPort(s.Host, strconv.Itoa(int(s.Port)))
}
