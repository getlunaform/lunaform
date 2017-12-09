package database

import (
	"testing"

	"reflect"
	"fmt"
)

var (
	dbTestType       = "test-type"
	dbTestKey        = "test-key"
	dbDuplicateKey   = "duplicate"
	dbNonexistantKey = "no-such-key"
	dbTestDoc        = map[string]string{"hello": "world"}
	dbTestDocU       = map[string]string{"jello": "whirled"}
)

func TestDatabaseInterface(t *testing.T) {

	// Test Redis bootstrapping and configure mock
	dbRedis, err := NewRedisDatabase("test", "localhost:3276", "", 0)
	t.Run("Redis DB does not error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}
	})
	dbRedis.client = &stubRedis{
		collections: make(map[string]interface{}),
	}

	// Test Memory Database bootstrap
	dbMem, err := NewMemoryDatabase()
	t.Run("Memory DB does not error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}
	})

	// Test JSON Database bootstrap
	f := mockFile{}
	dbJSON, err := NewJSONDatabase(f)
	t.Run("DB does not error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}
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
				if err != nil {
					t.Errorf("Unexpected error: %+v", err)
				}
			})

			t.Run("I get an error adding a collection which exists", func(t *testing.T) {
				err0 := db.Create(dbTestType, dbDuplicateKey, dbTestDoc)
				if err0 != nil {
					t.Errorf("Unexpected error: %+v", err0)
				}

				err1 := db.Create(dbTestType, dbDuplicateKey, dbTestDoc)
				if err1 == nil {
					t.Errorf("Expected an error:")
				}
			})

			var i map[string]string
			t.Run("I can read a collection", func(t *testing.T) {
				i = nil

				err := db.Read(dbTestType, dbTestKey, &i)
				if err != nil {
					t.Errorf("Unexpected error: %+v", err)
				}

				if !reflect.DeepEqual(i, dbTestDoc) {
					fmt.Printf("%T, %T", dbTestDoc, i)
					t.Errorf("expected %+v, received %+v", dbTestDoc, i)
				}
			})

			t.Run("I get an error reading a collection which doesn't exist", func(t *testing.T) {
				i = nil
				err := db.Read(dbTestType, dbNonexistantKey, &i)
				if err == nil {
					t.Errorf("Expected error")
				}

			})

			t.Run("I can update a collection", func(t *testing.T) {
				err := db.Update(dbTestType, dbTestKey, dbTestDocU)
				if err != nil {
					t.Errorf("Unexpected error: %+v", err)
				}

				i = nil
				err = db.Read(dbTestType, dbTestKey, &i)
				if err != nil {
					t.Errorf("Unexpected error: %+v", err)
				}

				if !reflect.DeepEqual(i, dbTestDocU) {
					t.Errorf("expected %+v, received %+v", dbTestDocU, i)
				}
			})

			t.Run("I get an error updating a collection which doesn't exist", func(*testing.T) {
				err := db.Update(dbTestType, dbNonexistantKey, dbTestDocU)
				if err == nil {
					t.Errorf("Expected error")
				}
			})

			t.Run("I can delete a collection", func(t *testing.T) {
				err := db.Delete(dbTestType, dbTestKey)
				if err != nil {
					t.Errorf("Unexpected error: %+v", err)
				}
			})
		})
	}
}
