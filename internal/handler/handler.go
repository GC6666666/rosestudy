package handler

import (
	"rose/common/net/chttp"
	"rose/internal/service"
)

var svc *service.Service

func InitRouter(s *chttp.Server, service *service.Service) {
	svc = service

	g := s.Group("/v1")
	ug := g.Group("/user")
	{
		ug.POST("/create", userCreate)
		ug.GET("/detail", userDetail)
		ug.GET("/list", userList)
	}
}
