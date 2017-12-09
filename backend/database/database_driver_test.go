package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dbTestType       = "test-type"
	dbTestKey        = "test-key"
	dbDuplicateKey   = "duplicate"
	dbNonexistantKey = "no-such-key"
	dbTestDoc        = map[string]string{"hello": "world"}
	dbTestDocU       = map[string]string{"jello": "whirled"}
	dbBadRecord      = map[chan int]string{make(chan int): "world"} // Intentially an unserialisable type
)

func TestDriverInterface(t *testing.T) {

	// Test Redis bootstrapping and configure mock
	dbRedis, err := NewRedisDatabase("test", "localhost:3276", "", 0)
	t.Run("Redis DB does not error", func(t *testing.T) {
		assert.NoError(t, err)
	})
	dbRedis.client = &stubRedis{
		collections: make(map[string]interface{}),
	}

	// Test Memory Database bootstrap
	dbMem, err := NewMemoryDatabase()
	t.Run("Memory DB does not error", func(t *testing.T) {
		assert.NoError(t, err)
	})

	// Test JSON Database bootstrap
	f := mockFile{}
	dbJSON, err := NewJSONDatabase(f)
	t.Run("DB does not error", func(t *testing.T) {
		assert.NoError(t, err)
	})

	for key, db := range map[string]Driver{
		"Redis":  dbRedis,
		"Memory": dbMem,
		"JSON":   dbJSON,
	} {
		t.Run(key, func(t *testing.T) {

			t.Run("I can ping my "+key, func(*testing.T) {
				_ = db.Ping()
			})

			t.Run("I can close my "+key, func(*testing.T) {
				_ = db.Close()
			})

			t.Run("I can add a collection", func(t *testing.T) {
				err := db.Create(dbTestType, dbTestKey, dbTestDoc)
				assert.NoError(t, err)
			})

			t.Run("I get an error adding a collection which exists", func(t *testing.T) {
				err0 := db.Create(dbTestType, dbDuplicateKey, dbTestDoc)
				assert.NoError(t, err0)

				err1 := db.Create(dbTestType, dbDuplicateKey, dbTestDoc)
				assert.Error(t, err1)
			})

			var i map[string]string
			t.Run("I can read a collection", func(t *testing.T) {
				i = nil

				err := db.Read(dbTestType, dbTestKey, &i)
				assert.NoError(t, err)
				assert.Equal(t, dbTestDoc, i)
			})

			t.Run("I get an error reading a collection which doesn't exist", func(t *testing.T) {
				i = nil
				err := db.Read(dbTestType, dbNonexistantKey, &i)
				assert.EqualError(t, err, "\"test-type\" \"no-such-key\" does not exist")
			})

			t.Run("I can update a collection", func(t *testing.T) {
				err := db.Update(dbTestType, dbTestKey, dbTestDocU)
				assert.NoError(t, err)

				i = nil
				err = db.Read(dbTestType, dbTestKey, &i)
				assert.NoError(t, err)

				assert.Equal(t, i, dbTestDocU)
			})

			t.Run("I get an error updating a collection which doesn't exist", func(*testing.T) {
				err := db.Update(dbTestType, dbNonexistantKey, dbTestDocU)
				assert.EqualError(t, err, "\"test-type\" \"no-such-key\" does not exist")
			})

			t.Run("I get an error creating a collection with a bad record key", func(*testing.T) {
				err := db.Update(dbTestType, dbTestKey, dbBadRecord)
				assert.EqualError(t, err, "json: unsupported type: map[chan int]string")
			})

			t.Run("I can delete a collection", func(t *testing.T) {
				err := db.Delete(dbTestType, dbTestKey)
				assert.NoError(t, err)
			})
		})
	}
}
