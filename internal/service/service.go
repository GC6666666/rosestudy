package service

import (
	"rose/internal/conf"
	"rose/internal/repo"
)

type Service struct {
	conf *conf.Conf
	repo *repo.Repo
}

func NewService(conf *conf.Conf) *Service {
	return &Service{
		conf: conf,
		repo: repo.NewRepo(conf),
	}
}
