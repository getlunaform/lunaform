package database

import (
	"github.com/zeebox/terraform-server/backend"
	"testing"
)

func TestJSONDB(t *testing.T) {
	var db backend.Driver
	var err error
	db, err = NewJSONDatabase("mock_path")
	t.Run("DB does not error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}
	})

	t.Run("I can ping my json", func(*testing.T) {
		_ = db.Ping()
	})

}
