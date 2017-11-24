package database

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

var (
	redTestType       = "test-type"
	redTestKey        = "test-key"
	redDuplicateKey   = "duplicate"
	redNonexistantKey = "no-such-key"
	redTestDoc        = map[string]string{"hello": "world"}
	redTestDocU       = map[string]string{"jello": "whirled"}
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
	keys := make([]string, len(r.collections))

	i := 0
	for k := range r.collections {
		keys[i] = k
		i++
	}

	return redis.NewStringSliceResult(keys, nil)
}

func (r *stubRedis) Get(s string) *redis.StringCmd {
	return redis.NewStringResult(r.collections[s].(string), nil)
}

func (r *stubRedis) Set(s string, i interface{}, t time.Duration) *redis.StatusCmd {
	v, _ := json.Marshal(i)
	r.collections[s] = string(v)

	return redis.NewStatusResult("", nil)
}

func TestRedisDatabase(t *testing.T) {
	db, err := NewRedisDatabase("test", "localhost:3276", "", 0)
	t.Run("DB does not error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}
	})

	db.client = &stubRedis{
		collections: make(map[string]interface{}),
	}

	t.Run("I can add a collection", func(t *testing.T) {
		err := db.Create(redTestType, redTestKey, redTestDoc)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}
	})

	t.Run("I get an error adding a collection which exists", func(t *testing.T) {
		err0 := db.Create(redTestType, redDuplicateKey, redTestDoc)
		if err0 != nil {
			t.Errorf("Unexpected error: %+v", err0)
		}

		err1 := db.Create(redTestType, redDuplicateKey, redTestDoc)
		if err1 == nil {
			t.Errorf("Expected an error:")
		}
	})

	var i map[string]string
	t.Run("I can read a collection", func(t *testing.T) {
		i = nil

		err := db.Read(redTestType, redTestKey, &i)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		if !reflect.DeepEqual(i, redTestDoc) {
			fmt.Printf("%T, %T", redTestDoc, i)
			t.Errorf("expected %+v, received %+v", redTestDoc, i)
		}
	})

	t.Run("I get an error reading a collection which doesn't exist", func(t *testing.T) {
		i = nil
		err := db.Read(redTestType, redNonexistantKey, &i)
		if err == nil {
			t.Errorf("Expected error")
		}

	})

	t.Run("I can update a collection", func(t *testing.T) {
		err := db.Update(redTestType, redTestKey, redTestDocU)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		i = nil
		err = db.Read(redTestType, redTestKey, &i)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		if !reflect.DeepEqual(i, redTestDocU) {
			t.Errorf("expected %+v, received %+v", redTestDocU, i)
		}
	})

	t.Run("I get an error updating a collection which doesn't exist", func(*testing.T) {
		err := db.Update(redTestType, redNonexistantKey, redTestDocU)
		if err == nil {
			t.Errorf("Expected error")
		}
	})

	t.Run("I can delete a collection", func(t *testing.T) {
		err := db.Delete(redTestType, redTestKey)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}
	})
}
