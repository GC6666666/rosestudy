package chttp

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"sync/atomic"
)

type Server struct {
	*gin.RouterGroup

	conf   *Config
	server atomic.Value
	engine *gin.Engine
}

func NewServer(conf *Config) *Server {
	s := &Server{
		conf:   conf,
		engine: gin.New(),
	}

	s.RouterGroup = &s.engine.RouterGroup

	return s
}

func (s *Server) Start() error {
	lis, err := net.Listen(s.conf.Network, s.conf.Address)
	if err != nil {
		return err
	}

	hs := &http.Server{
		Handler:      s.engine,
		ReadTimeout:  s.conf.ReadTimeOut,
		WriteTimeout: s.conf.WriteTimeOut,
	}
	s.server.Store(hs)

	return hs.Serve(lis)
}

func (s *Server) getServer() *http.Server {
	// atomic.Value存的是interface，进行类型断言，判断是不是http.Server，是的话返回true
	server, ok := s.server.Load().(*http.Server)
	if !ok {
		return nil
	}

	return server
}

func (s *Server) ShutDown(ctx context.Context) error {
	server := s.getServer()
	if server == nil {
		return fmt.Errorf("[chttp] server is nil")
	}

	return server.Shutdown(ctx)
}
