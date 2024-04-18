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
	/*
	   这里把gin的Engine作为http server的Handler来使用，Handler是一个接口，也就是gin的Engine实现了ServeHTTP方法，
	   当有http请求过来的时候都会走到Engine的ServeHTTP方法。
	*/
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
	// 这里获取server的作用是什么
	server := s.getServer()
	if server == nil {
		return fmt.Errorf("[chttp] server is nil")
	}

	return server.Shutdown(ctx)
}
