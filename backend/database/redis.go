package database

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	redisDelimeter = "."
)

// RedisDatabase represents a redis store for our server
// It holds state and docs and all kinds of awesome sauce
type RedisDatabase struct {
	client    redisClient
	namespace string
}

type redisClient interface {
	Close() error
	Del(...string) *redis.IntCmd
	Keys(string) *redis.StringSliceCmd
	Get(string) *redis.StringCmd
	Ping() *redis.StatusCmd
	Set(string, interface{}, time.Duration) *redis.StatusCmd
}

// NewRedisDatabase returns a redis database object
func NewRedisDatabase(namespace, address, password string, database int) (r RedisDatabase, err error) {
	r = RedisDatabase{
		client: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       database,
		}),
		namespace: namespace,
	}

	return
}

// Ping checks whether redis is up and usable
// it returns an error to determine this; a nil value means the redis database is usable
func (r RedisDatabase) Ping() error {
	_, err := r.client.Ping().Result()

	return err
}

// Close will close a connection to the redis database
func (r RedisDatabase) Close() error {
	return r.client.Close()
}

// Create a record within redis
func (r RedisDatabase) Create(recordType, key string, doc interface{}) error {
	k := r.key(recordType, key)

	if r.exists(k) {
		return fmt.Errorf("%q %q already exists at %q", recordType, key, k)
	}

	return r.set(k, doc)
}

// Read a record from redis
func (r RedisDatabase) Read(recordType, key string, i interface{}) (err error) {
	k := r.key(recordType, key)

	if !r.exists(k) {
		err = fmt.Errorf("%q %q does not exist", recordType, key)

		return
	}

	d := r.client.Get(k)
	if d.Err() != nil {
		err = d.Err()

		return
	}

	result, err := d.Result()

	return r.deserialize(result, i)
}

// Update a redis record
func (r RedisDatabase) Update(recordType, key string, doc interface{}) (err error) {
	k := r.key(recordType, key)

	if !r.exists(k) {
		return fmt.Errorf("%q %q does not exist", recordType, key)
	}

	return r.set(k, doc)
}

// Delete a record from redis
func (r RedisDatabase) Delete(recordType, key string) error {
	k := r.key(recordType, key)

	return r.client.Del(k).Err()
}

func (r RedisDatabase) deserialize(s string, i interface{}) (err error) {
	cleanString, err := strconv.Unquote(s)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(cleanString), i)

	return
}

func (r RedisDatabase) exists(key string) bool {
	k, _ := r.client.Keys(key).Result()

	return len(k) > 0
}

func (r RedisDatabase) serialize(i interface{}) (string, error) {
	v, err := json.Marshal(i)

	return string(v), err
}

func (r RedisDatabase) set(key string, doc interface{}) error {
	v, err := r.serialize(doc)
	if err != nil {
		return err
	}

	return r.client.Set(key, v, 0).Err()
}

func (r RedisDatabase) key(recordType, key string) string {
	return fmt.Sprintf("%s%s%s%s%s", r.namespace, redisDelimeter, recordType, redisDelimeter, key)
}
