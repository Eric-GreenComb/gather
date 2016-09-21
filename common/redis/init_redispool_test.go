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

	_testRedis := "banerwai:gather:common:redis:test"
	_testValue := "redis_test"
	_, err := conn.Do("SET", _testRedis, _testValue)

	if err != nil {
		t.Fatalf("SET err : %v", err)
	}

	_value, err := redis.String(conn.Do("GET", _testRedis))

	if err != nil {
		t.Fatalf("GET err : %v", err)
	}

	if _value != _testValue {
		t.Fatalf("GET data err")
	}

	_, err = conn.Do("LPUSH", "banerwai:email", "1")
	_, err = conn.Do("LPUSH", "banerwai:email", "2")
	_, err = conn.Do("LPUSH", "banerwai:email", "3")

	var _errPop error
	var _email string
	for _errPop == nil {
		_email, _errPop = redis.String(conn.Do("LPOP", "banerwai:email"))
		if _errPop != nil {
			fmt.Println(_errPop.Error())
			continue
		}
		fmt.Println("email", _email)
	}
}
