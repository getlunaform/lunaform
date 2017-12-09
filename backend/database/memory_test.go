package database

import (
	"github.com/zeebox/terraform-server/backend"
	"reflect"
	"testing"
)

var (
	memTestType       = "test-type"
	memTestKey        = "test-key"
	memDuplicateKey   = "duplicate"
	memNonExistantKey = "no-such-key"
	memTestDoc        = map[string]string{"hello": "world"}
	memTestDocU       = map[string]string{"jello": "whirled"}
)

func TestMemoryDB(t *testing.T) {
	db, err := NewMemoryDatabase()

	t.Run("DB does not error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}
	})

	t.Run("I can use redis as a DB interface", func(*testing.T) {
		var _ backend.Driver = (*MemoryDatabase)(nil)
	})

	t.Run("I can ping my memory database", func(*testing.T) {
		_ = db.Ping()
	})

	t.Run("I can close my memory database", func(*testing.T) {
		_ = db.Close()
	})

	t.Run("I can add a collection", func(t *testing.T) {
		l0 := len(db.collections)

		err := db.Create(memTestType, memTestKey, memTestDoc)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		l1 := len(db.collections)
		if l1 != l0+1 {
			t.Errorf("Create did not add to collections. Contents are: %+v", db.collections)
		}
	})

	t.Run("I receive an error adding a collection which exists", func(t *testing.T) {
		err := db.Create(memTestType, memDuplicateKey, memTestDoc)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		err = db.Create(memTestType, memDuplicateKey, memTestDoc)
		if err == nil {
			t.Errorf("Expected error")
		}
	})

	var i map[string]string
	t.Run("I can read a collection", func(t *testing.T) {
		i = nil

		err := db.Read(memTestType, memTestKey, &i)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		if !reflect.DeepEqual(i, memTestDoc) {
			t.Errorf("expected %+v, received %+v", memTestDoc, i)
		}
	})

	t.Run("I get an error reading a collection which does not exist", func(t *testing.T) {
		i = nil

		err := db.Read(memTestType, memNonExistantKey, &i)
		if err == nil {
			t.Errorf("Expected error")
		}
	})

	t.Run("I can update a collection", func(t *testing.T) {
		l0 := len(db.collections)

		err := db.Update(memTestType, memTestKey, memTestDocU)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		l1 := len(db.collections)
		if l0 != l1 {
			t.Errorf("Update did not update collections. Contents are: %+v", db.collections)
		}

		i := make(map[string]string)
		err = db.Read(memTestType, memTestKey, &i)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		if !reflect.DeepEqual(i, memTestDocU) {
			t.Errorf("expected %+v, received %+v", memTestDoc, i)
		}
	})

	t.Run("I get an error updating a collection which does not exist", func(t *testing.T) {
		err := db.Update(memTestType, memNonExistantKey, memTestDocU)
		if err == nil {
			t.Errorf("Expected error")
		}
	})

	t.Run("I can delete a collection", func(t *testing.T) {
		l0 := len(db.collections)

		err := db.Delete(memTestType, memTestKey)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		l1 := len(db.collections)
		if l1 != l0-1 {
			t.Errorf("Delete did not remove a collection. Contents are: %+v", db.collections)
		}

	})
}
