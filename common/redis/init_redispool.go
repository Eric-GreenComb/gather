package redis

import (
	"github.com/garyburd/redigo/redis"

	"github.com/banerwai/global/constant"
	"github.com/banerwai/gommon/etcd"
)

// 重写生成连接池方法
func newPool() *redis.Pool {
	_addr, _err := etcd.GetValue(constant.EtcdKeyRedisAddr)
	if _err != nil {
		return nil
	}
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", _addr)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// RedisPool 生成连接池
var RedisPool = newPool()
