package biz

import (
	"eshop/app/user/config"
	"eshop/app/user/internal/data"
	"sync"
)

type BusinessHandler struct {
	data *data.Data
	conf *config.Config
	sync.Mutex
	initialized bool
}

func New(conf *config.Config) *BusinessHandler {
	b := new(BusinessHandler)
	b.conf = conf
	return b
}

func (b *BusinessHandler) Initialized() error {
	b.Lock()
	defer b.Unlock()
	if b.initialized {
		return nil
	}
	b.data = data.New(b.conf)
	if err := b.data.Initialize(); err != nil {
		return err
	}
	b.initialized = true
	return nil
}
