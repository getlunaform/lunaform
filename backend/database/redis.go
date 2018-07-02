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

// redisDatabase represents a redis store for our server
// It holds state and docs and all kinds of awesome sauce
type redisDatabase struct {
	client    RedisClient
	namespace string
}

// RedisClient interface to allow custom redis clients to be used
type RedisClient interface {
	Close() error
	Del(...string) *redis.IntCmd
	Keys(string) *redis.StringSliceCmd
	Get(string) *redis.StringCmd
	Ping() *redis.StatusCmd
	Set(string, interface{}, time.Duration) *redis.StatusCmd
}

// NewRedisDBDriver returns a redis database object
func NewRedisDBDriver(namespace, address, password string, database int, driver RedisClient) (r Driver, err error) {
	if driver == nil {
		driver = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       database,
		})
	}
	r = redisDatabase{
		client:    driver,
		namespace: namespace,
	}

	return
}

// Ping checks whether redis is up and usable
// it returns an error to determine this; a nil value means the redis database is usable
func (r redisDatabase) Ping() error {
	_, err := r.client.Ping().Result()

	return err
}

// Close will close a connection to the redis database
func (r redisDatabase) Close() error {
	return r.client.Close()
}

// Create a record within redis
func (r redisDatabase) Create(recordType, key string, doc interface{}) error {
	k := r.key(recordType, key)

	if r.exists(k) {
		return fmt.Errorf("%q %q already exists at %q", recordType, key, k)
	}

	return r.set(k, doc)
}

// Read a record from redis
func (r redisDatabase) Read(recordType, key string, i interface{}) (err error) {
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
func (r redisDatabase) Update(recordType, key string, doc interface{}) (err error) {
	k := r.key(recordType, key)

	if !r.exists(k) {
		return fmt.Errorf("%q %q does not exist", recordType, key)
	}

	return r.set(k, doc)
}

// Delete a record from redis
func (r redisDatabase) Delete(recordType, key string) error {
	k := r.key(recordType, key)

	return r.client.Del(k).Err()
}

func (r redisDatabase) List(recordType string) (err error) {
	return nil
}

func (r redisDatabase) deserialize(s string, i interface{}) (err error) {
	cleanString, err := strconv.Unquote(s)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(cleanString), i)

	return
}

func (r redisDatabase) exists(key string) bool {
	k, _ := r.client.Keys(key).Result()

	return len(k) > 0
}

func (r redisDatabase) serialize(i interface{}) (string, error) {
	v, err := json.Marshal(i)

	return string(v), err
}

func (r redisDatabase) set(key string, doc interface{}) error {
	v, err := r.serialize(doc)
	if err != nil {
		return err
	}

	return r.client.Set(key, v, 0).Err()
}

func (r redisDatabase) key(recordType, key string) string {
	return fmt.Sprintf("%s%s%s%s%s", r.namespace, redisDelimeter, recordType, redisDelimeter, key)
}
