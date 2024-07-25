package limiter

import "time"

//go:generate mockery --name Store --outpkg mock --output mock --filename store.go --with-expecter=true

type Store interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
	Incr(key string) error
	Expire(key string, expiration time.Duration) error
}
