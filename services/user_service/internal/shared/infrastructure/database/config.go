package database

import (
	"errors"
	"fmt"

	"github.com/caarlos0/env/v11"
)

type DBConfig struct {
	Host string `env:"DB_HOST,required"`
	User string `env:"DB_USER,required"`
	Pass string `env:"DB_PASS,required"`
	Name string `env:"DB_NAME,required"`
}

func NewDBConfig() (*DBConfig, error) {
	var cfg DBConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New(fmt.Sprintf("環境変数が設定されていません: %v", err))
	}
	return &cfg, nil
}

func (d DBConfig) GetHost() string {
	return d.Host
}

func (d DBConfig) GetUser() string {
	return d.User
}

func (d DBConfig) GetPass() string {
	return d.Pass
}

func (d DBConfig) GetName() string {
	return d.Name
}
