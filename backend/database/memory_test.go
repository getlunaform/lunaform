package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeebox/terraform-server/backend"
)

func TestMemoryDB(t *testing.T) {
	t.Run("Interface", testMemoryDBInterface)
	t.Run("Insert", testMemoryDBInsert)
}

func testMemoryDBInterface(t *testing.T) {
	var db backend.Database
	db = MemoryDatabase{}
	assert.NotNil(t, db)
}

func testMemoryDBInsert(t *testing.T) {
	var db backend.Database

	db = NewMemoryDatabase()
	err := db.Insert(
		"test",
		backend.NewCollection("test", db),
	)
	assert.Nil(t, err)

	col := db.Collection("test")
	assert.NotNil(t, col)

	assert.Equal(t, col.Name(), "test")
	assert.Equal(t, col.Count(), 0)

	err = col.Create("test-key", "test-value")
	assert.Nil(t, err)

	var result string
	err = col.Read("test-key", &result)
	assert.Equal(t, "test-value", result)

	var result1 string
	col.Delete("test-key")
	err = col.Read("test-key", &result1)
	assert.Empty(t, result1)

	db.Close()
}
