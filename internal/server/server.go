package server

import (
	"rose/common/net/chttp"
	"rose/internal/conf"
	"rose/internal/handler"
	"rose/internal/service"
)

func NewHTTP(conf *conf.Conf) *chttp.Server {
	s := chttp.NewServer(conf.Server)
	svc := service.NewService(conf)

	handler.InitRouter(s, svc)

	err := s.Start()
	if err != nil {
		panic(err)
	}

	return s
}
