package redis

import (
	"fmt"
	"testing"

	"github.com/garyburd/redigo/redis"
)

// need start micro render service localhost:6003
func TestRedisPool(t *testing.T) {

	conn := RedisPool.Get()
	defer conn.Close()

	_redis_test := "banerwai:gather:common:redis:test"
	_test_value := "redis_test"
	_, err := conn.Do("SET", _redis_test, _test_value)

	if err != nil {
		t.Fatalf("SET err : %v", err)
	}

	_value, err := redis.String(conn.Do("GET", _redis_test))

	if err != nil {
		t.Fatalf("GET err : %v", err)
	}

	if _value != _test_value {
		t.Fatalf("GET data err")
	}

	_, err = conn.Do("LPUSH", "banerwai:email", "1")
	_, err = conn.Do("LPUSH", "banerwai:email", "2")
	_, err = conn.Do("LPUSH", "banerwai:email", "3")

	var _pop_err error
	var _email string
	for _pop_err == nil {
		_email, _pop_err = redis.String(conn.Do("LPOP", "banerwai:email"))
		if _pop_err != nil {
			fmt.Println(_pop_err.Error())
			continue
		}
		fmt.Println("email", _email)
	}

}
