package data

import (
	user "eshop/api/user/v1"
	"eshop/app/user/config"
	"eshop/app/user/internal/model"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
	"time"
)

type Data struct {
	Db         *xorm.Engine
	Cache      *redis.Client
	initialize bool
	sync.Mutex
	conf *config.Config
}

func New(conf *config.Config) *Data {
	d := new(Data)
	d.conf = conf
	d.initialize = false
	return d
}

func (d *Data) Initialize() {
	d.Lock()
	defer d.Unlock()
	if d.initialize {
		return
	}
	d.initDB(d.conf.Db)
	d.initCache(d.conf.Redis)
	d.initialize = true
}

func (d *Data) initDB(config *config.DbConfig) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.User, config.Passwd, config.Host, config.Port, config.DBName)
	engine, err := xorm.NewEngine("mysql", uri)
	if err != nil {
		log.Fatalf("init db failed,%+v", err)
	}
	d.Db = engine
	d.Db.SetMaxOpenConns(config.MaxConn)
	d.Db.SetMaxIdleConns(config.MaxIdle)
	d.Db.SetConnMaxLifetime(time.Duration(config.MaxLifeTime))
	d.Db.ShowSQL(config.ShowLog)
	err = d.Db.Ping()
	if err != nil {
		log.Fatalf("init db failed,err:%v", err)
	}
	err = d.Db.Sync2(new(model.User), new(model.Token))
	if err != nil {
		log.Fatalf("init db failed,err:%v", err)
	}
}

func (d *Data) initCache(config *config.RedisConfig) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Passwd,
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("init cache failed,err:%v", err)
	}
	d.Cache = client
	log.Println(pong)
}

func (d *Data) Page(p *user.Page) (int, int) {
	if p == nil {
		return 0, 10
	}
	if p.Page == 0 {
		p.Page = 1
	}
	limit := 10
	if p.Limit > 0 {
		limit = int(p.Limit)
	}
	return (int(p.Page) - 1) * limit, limit
}

type Operation func(session *xorm.Session) error

func (d *Data) Transaction(session *xorm.Session, f Operation) error {
	var err error
	if err = session.Begin(); err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			session.Rollback()
			panic(p)
		} else if err != nil {
			session.Rollback()
		} else {
			err = session.Commit()
		}
	}()
	err = f(session)
	return err
}
