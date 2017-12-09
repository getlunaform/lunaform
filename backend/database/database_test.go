package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBFactory(t *testing.T) {
	db := NewDatabaseWithDriver(memoryDatabase{})
	assert.NotNil(t, db.driver)
}

func TestDBRecord(t *testing.T) {
	for _, test := range []struct {
		Key  string
		Type string
	}{
		{Key: "hello", Type: "string"},
	} {

		m := make(map[string]interface{}, 2)
		m["Key"] = test.Key
		m["Type"] = test.Type

		r := Record(m)

		assert.Equal(t, r.Key(), test.Key)
		assert.Equal(t, r.Type(), test.Type)
	}

}
