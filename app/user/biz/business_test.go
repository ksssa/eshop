package biz

import (
	"eshop/app/user/config"
	"os"
	"testing"
)

var bizHandler *BusinessHandler

func TestMain(m *testing.M) {
	conf := &config.Config{
		Db: &config.DbConfig{
			Host:        "192.168.1.10",
			Port:        3306,
			User:        "zhu",
			Passwd:      "111111",
			Prefix:      "es_",
			DBName:      "eshop",
			MaxConn:     10,
			MaxIdle:     5,
			MaxLifeTime: 600,
			ShowLog:     true,
		},
		Redis: &config.RedisConfig{
			Host:   "192.168.1.10",
			Port:   6379,
			Passwd: "111111",
		},
		Etcd: nil,
	}
	bizHandler = New(conf)
	bizHandler.Initialized()
	exitVal := m.Run()
	os.Exit(exitVal)
}
