package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Config struct {
	DSN         string        `yaml:"dsn"`
	Active      int           `yaml:"active"`
	Idle        int           `yaml:"idle"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

type DB struct {
	*sql.DB
}

func NewDB(conf *Config) *DB {
	if conf == nil {
		panic("conf cannot be nil")
	}

	d, err := sql.Open("mysql", conf.DSN)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(conf.Active)
	d.SetMaxIdleConns(conf.Idle)
	d.SetConnMaxIdleTime(conf.IdleTimeout)

	return &DB{
		DB: d,
	}
}
