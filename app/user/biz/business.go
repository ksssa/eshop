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

func (b *BusinessHandler) Initialized() {
	b.Lock()
	defer b.Unlock()
	if b.initialized {
		return
	}
	b.data.Initialize(b.conf)
	b.initialized = true
}
