package config

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
	Db    *DbConfig    `yaml:"db"`
	Redis *RedisConfig `yaml:"redis"`
	Etcd  *EtcdConfig  `yaml:"etcd"`
}
