package innerr

import "errors"

var (
	ErrDataNotExist  = errors.New("data not exist")
	ErrCacheNotExist = errors.New("redis: nil")
	ErrTokenExpire   = errors.New("access token is expire")
)
