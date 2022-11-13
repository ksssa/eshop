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
		Host:   "192.168.1.10",
		Port:   6379,
		Passwd: "111111",
	}

	d.initCache(cacheCfg)
	val := m.Run()
	os.Exit(val)
}

func TestRedis(t *testing.T) {
	res, err := d.Cache.Set("eshop", "1111", time.Duration(time.Now().Unix())).Result()
	require.NoError(t, err)
	fmt.Printf("%+v\n", res)
	res, err = d.Cache.Get("eshop1").Result()
	fmt.Println(res, err)
	require.NoError(t, err)
	fmt.Println(res)
}
