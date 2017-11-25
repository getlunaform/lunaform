package database

import "testing"

func TestJSONDB(t *testing.T) {
	db, err := NewJSONDatabase()
	t.Run("DB does not error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}
	})
}
