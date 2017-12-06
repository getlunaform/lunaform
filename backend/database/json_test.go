package database

import (
	"github.com/zeebox/terraform-server/backend"
	"io"
	"os"
	"reflect"
	"testing"
)

var (
	jsonTestType       = "test-type"
	jsonTestKey        = "test-key"
	jsonDuplicateKey   = "duplicate"
	jsonNonExistantKey = "no-such-key"
	jsonTestDoc        = map[string]string{"hello": "world"}
	jsonTestDocU       = map[string]string{"jello": "whirled"}
)

type mockCall struct {
	Name      string
	Arguments []interface{}
}

type mockFile struct {
	Calls []mockCall
}

func (f mockFile) Close() error {
	f.Calls = append(f.Calls, mockCall{
		Name: "Close",
	})
	return nil
}

func (f mockFile) Read(b []byte) (int, error) {
	f.Calls = append(f.Calls, mockCall{
		Name:      "Read",
		Arguments: []interface{}{b},
	})
	r := "{ \"collections 0 name\": \"hello\"}"
	copy(b[:], []byte(r))
	return len(r), io.EOF
}

func (f mockFile) Seek(a int64, b int) (int64, error) {
	f.Calls = append(f.Calls, mockCall{
		Name:      "Seek",
		Arguments: []interface{}{a, b},
	})
	return 0, nil
}

func (f mockFile) Stat() (os.FileInfo, error) {
	f.Calls = append(f.Calls, mockCall{
		Name: "Stat",
	})
	return nil, nil
}

func (f mockFile) Write(b []byte) (int, error) {
	f.Calls = append(f.Calls, mockCall{
		Name:      "Write",
		Arguments: []interface{}{b},
	})
	return 0, nil
}

func (f mockFile) WriteAt(b []byte, i int64) (int, error) {
	f.Calls = append(f.Calls, mockCall{
		Name:      "WriteAt",
		Arguments: []interface{}{b},
	})
	return 0, nil
}

func (f mockFile) Truncate(i int64) (err error) {
	f.Calls = append(f.Calls, mockCall{
		Name:      "Write",
		Arguments: []interface{}{i},
	})
	return nil
}

func TestJSONDB(t *testing.T) {
	f := mockFile{}
	db, err := NewJSONDatabase(f)
	t.Run("DB does not error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}
	})

	t.Run("I can use json as a DB interface", func(*testing.T) {
		var _ backend.Driver = (*JSONDatabase)(nil)
	})

	t.Run("I can use a file in the JSONDatabase", func(*testing.T) {
		var db1 JSONDatabase
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}
		db1, err = NewJSONDatabase(f)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}
		db1.Close()
	})

	t.Run("I can ping my json database", func(*testing.T) {
		_ = db.Ping()
	})

	t.Run("I can close my json database", func(*testing.T) {
		_ = db.Close()
	})

	t.Run("I can add a collection", func(t *testing.T) {
		l0 := len(db.db.collections)

		err := db.Create(jsonTestType, jsonTestKey, jsonTestDoc)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		l1 := len(db.db.collections)
		if l1 != l0+1 {
			t.Errorf("Create did not add to collections. Contents are: %+v", db.db.collections)
		}
	})

	t.Run("I receive an error adding a collection which exists", func(t *testing.T) {
		err := db.Create(jsonTestType, jsonDuplicateKey, jsonTestDoc)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		err = db.Create(jsonTestType, jsonDuplicateKey, jsonTestDoc)
		if err == nil {
			t.Errorf("Expected error")
		}
	})

	var i map[string]string
	t.Run("I can read a collection", func(t *testing.T) {
		i = nil

		err := db.Read(jsonTestType, jsonTestKey, &i)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		if !reflect.DeepEqual(i, jsonTestDoc) {
			t.Errorf("expected %+v, received %+v", jsonTestDoc, i)
		}
	})

	t.Run("I get an error reading a collection which does not exist", func(t *testing.T) {
		i = nil

		err := db.Read(jsonTestType, jsonNonExistantKey, &i)
		if err == nil {
			t.Errorf("Expected error")
		}
	})

	t.Run("I can update a collection", func(t *testing.T) {
		l0 := len(db.db.collections)

		err := db.Update(jsonTestType, jsonTestKey, jsonTestDocU)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		l1 := len(db.db.collections)
		if l0 != l1 {
			t.Errorf("Update did not update collections. Contents are: %+v", db.db.collections)
		}

		i := make(map[string]string)
		err = db.Read(jsonTestType, jsonTestKey, &i)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		if !reflect.DeepEqual(i, jsonTestDocU) {
			t.Errorf("expected %+v, received %+v", jsonTestDoc, i)
		}
	})

	t.Run("I get an error updating a collection which does not exist", func(t *testing.T) {
		err := db.Update(jsonTestType, jsonNonExistantKey, jsonTestDocU)
		if err == nil {
			t.Errorf("Expected error")
		}
	})

	t.Run("I can delete a collection", func(t *testing.T) {
		l0 := len(db.db.collections)

		err := db.Delete(jsonTestType, jsonTestKey)
		if err != nil {
			t.Errorf("Unexpected error: %+v", err)
		}

		l1 := len(db.db.collections)
		if l1 != l0-1 {
			t.Errorf("Delete did not remove a collection. Contents are: %+v", db.db.collections)
		}

	})
