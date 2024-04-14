package chttp

import "time"

type Config struct {
	Network      string        `yaml:"network"`
	Address      string        `yaml:"address"`
	ReadTimeOut  time.Duration `yaml:"readTimeOut"`
	WriteTimeOut time.Duration `yaml:"writeTimeOut"`
}
