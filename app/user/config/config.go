package config

import (
	"github.com/go-kratos/kratos/v2/log"
	"gopkg.in/yaml.v2"
	"os"
)

type ServiceConfig struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	LogLevel string `yaml:"logLevel"`
	LogPath  string `yaml:"logPath"`
}

type DbConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Passwd      string `yaml:"passwd"`
	Prefix      string `yaml:"prefix"`
	DBName      string `yaml:"DBName"`
	MaxConn     int    `yaml:"maxConn"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxLifeTime int64  `yaml:"maxLifeTime"`
	ShowLog     bool   `yaml:"showLog"`
}

type RedisConfig struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Passwd string `yaml:"passwd"`
}

type EtcdConfig struct {
}

type Config struct {
	Db      *DbConfig      `yaml:"db"`
	Redis   *RedisConfig   `yaml:"redis"`
	Etcd    *EtcdConfig    `yaml:"etcd"`
	Service *ServiceConfig `yaml:"service"`
}

func LoadConfig(path string) (*Config, error) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	c := new(Config)
	err = yaml.Unmarshal(configFile, c)
	return c, err
}

func CheckConfig(c *Config) bool {
	if c == nil || c.Db == nil || c.Service == nil || c.Redis == nil {
		log.Error("necessary config is nil")
		return false
	}
	if len(c.Service.Host) == 0 {
		log.Error("host is empty")
		return false
	}
	if len(c.Db.Host) == 0 {
		log.Error("db host is empty")
		return false
	}
	return true
}
