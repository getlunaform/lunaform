package database

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"strings"
)

type stubRedis struct {
	collections map[string]interface{}
}

func (r *stubRedis) Close() error { return nil }

func (r *stubRedis) Ping() *redis.StatusCmd {
	return redis.NewStatusResult("", nil)
}

func (r *stubRedis) Del(s ...string) *redis.IntCmd {
	c := 0
	for _, k := range s {
		c++
		delete(r.collections, k)
	}

	return redis.NewIntResult(1, nil)
}

func (r *stubRedis) Keys(s string) *redis.StringSliceCmd {
	keys := []string{}

	for k := range r.collections {
		if s == "*" || k == s {
			keys = append(keys, k)
		} else if strings.HasSuffix(s, "*") {
			prefix := s[0 : len(s)-1]
			if strings.HasPrefix(k, prefix) {
				keys = append(keys, k)
			}
		}
	}

	return redis.NewStringSliceResult(keys, nil)
}

func (r *stubRedis) Get(s string) *redis.StringCmd {
	return redis.NewStringResult(r.collections[s].(string), nil)
}

func (r *stubRedis) MGet(keys ...string) *redis.SliceCmd {
	res := make([]interface{}, len(keys))
	for i, key := range keys {
		res[i] = r.collections[key]
	}
	return redis.NewSliceResult(res, nil)
}

func (r *stubRedis) Set(s string, i interface{}, t time.Duration) *redis.StatusCmd {
	v, _ := json.Marshal(i)
	r.collections[s] = string(v)

	return redis.NewStatusResult("", nil)
}
