package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBFactory(t *testing.T) {
	db := NewDatabaseWithDriver(&memoryDatabase{})
	assert.NotNil(t, db.driver)
}
