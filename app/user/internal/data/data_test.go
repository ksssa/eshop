package data

import (
	"eshop/app/user/config"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

var d *Data

func TestMain(m *testing.M) {

	d = new(Data)

	cacheCfg := &config.RedisConfig{
		Host:   "10.0.2.222",
		Port:   6379,
		Passwd: "111111",
	}

	d.initCache(cacheCfg)
	val := m.Run()
	os.Exit(val)
}

func TestRedis(t *testing.T) {
	res, err := d.cache.Set("eshop", "1111", time.Duration(time.Now().Unix())).Result()
	require.NoError(t, err)
	fmt.Printf("%+v", res)
}
